package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

type DBPool struct {
	mu       sync.Mutex
	pool     chan *sql.DB
	maxConns int
	dsn      string
}

func NewDBPool(dsn string, maxConns int) (*DBPool, error) {
	pool := make(chan *sql.DB, maxConns)

	for i := 0; i < maxConns; i++ {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxLifetime(5 * time.Minute)
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(maxConns)

		pool <- db
	}

	return &DBPool{
		pool:     pool,
		maxConns: maxConns,
		dsn:      dsn,
	}, nil
}

func (p *DBPool) Get() (*sql.DB, error) {
	select {
	case db := <-p.pool:
		return db, nil
	default:
		db, err := sql.Open("mysql", p.dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}

func (p *DBPool) Put(db *sql.DB) {
	select {
	case p.pool <- db:
	default:
		_ = db.Close()
	}
}

func (p *DBPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	close(p.pool)

	for db := range p.pool {
		_ = db.Close()
	}
}

func main() {
	dsn := "your_user:your_password@unix(/tmp/mysql.sock)/your_database"


	pool, err := NewDBPool(dsn, 5)
	if err != nil {
		log.Fatalf("Failed to initialize DB pool: %v", err)
	}

	db, err := pool.Get()
	if err != nil {
		log.Fatalf("Failed to get a connection: %v", err)
	}

	var now string
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Println("Current time:", now)

	pool.Put(db)

	pool.Close()
}
