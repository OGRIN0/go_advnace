package main 

import (
	"sync"
)

type Collection struct {
	Mutex sync.Mutex 
	Data map[string]string 
}

func NewCollection() Collection {
	return Collection{
		Data: make(map[string]string),
	}
}

func (c *Collection) Has(key string) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if _, ok := c.Data[key]; ok {
		return true 
	}
	return false 
}

func (c *Collection) Add(key, value string){
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if _, ok := c.Data[key]; ok {
		return 
	}
	c.Data[key] = value 
}

func main() {
	c := NewCollection()
	c.Add("a", "cake")
}
