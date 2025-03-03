package main

import (
	"fmt"
	"sync"
)

type Collection struct {
	RWMutex sync.RWMutex
	Data    map[string]string
}

func NewCollection() *Collection {
	return &Collection{
		Data: make(map[string]string),
	}
}

func (c *Collection) Has(key string) bool {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	_, ok := c.Data[key]
	return ok
}

func (c *Collection) Add(key, value string) {
	// First, check if the key exists without acquiring a write lock
	if c.Has(key) {
		return
	}

	// Now, acquire a write lock to add data
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.Data[key] = value
}

func main() {
	c := NewCollection()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.Add("a", "apple")
	}()

	go func() {
		defer wg.Done()
		c.Add("b", "banana")
	}()

	wg.Wait()
	fmt.Println("Execution completed")
}
