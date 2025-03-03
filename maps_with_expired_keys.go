package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration int64
}

type ExpiringMap struct {
	mutex sync.Mutex
	store map[string]item
}

func NewExpiringMap(cleanupInterval time.Duration) *ExpiringMap {
	em := &ExpiringMap{
		store: make(map[string]item),
	}
	go em.cleanupExpiredKeys(cleanupInterval)
	return em
}

func (em *ExpiringMap) Set(key string, value interface{}, duration time.Duration) {
	em.mutex.Lock()
	defer em.mutex.Unlock()
	em.store[key] = item{
		value:      value,
		expiration: time.Now().Add(duration).UnixNano(),
	}
}

func (em *ExpiringMap) Get(key string) (interface{}, bool) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	item, found := em.store[key]
	if !found || time.Now().UnixNano() > item.expiration {
		return nil, false
	}
	return item.value, true
}

func (em *ExpiringMap) cleanupExpiredKeys(interval time.Duration) {
	for {
		time.Sleep(interval)
		em.mutex.Lock()
		for key, item := range em.store {
			if time.Now().UnixNano() > item.expiration {
				delete(em.store, key)
			}
		}
		em.mutex.Unlock()
	}
}

func main() {
	em := NewExpiringMap(2 * time.Second)

	em.Set("foo", "bar", 3*time.Second)
	fmt.Println("Set key: foo -> bar (expires in 3s)")

	time.Sleep(2 * time.Second)
	if value, found := em.Get("foo"); found {
		fmt.Println("Found foo:", value)
	} else {
		fmt.Println("Key foo expired")
	}

	time.Sleep(2 * time.Second)
	if value, found := em.Get("foo"); found {
		fmt.Println("Found foo:", value)
	} else {
		fmt.Println("Key foo expired")
	}
}
