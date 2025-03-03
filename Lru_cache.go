package main 

import (
	"fmt"
	"container/list"
)

type LRUCache struct {
	capacity   int 
	cache      map[int]*list.Element 
	ll         *list.List 
}

type entry struct {
	key   int
	value int 
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (lru *LRUCache) Get(key int) int{
	if elem, found := lru.cache[key]; found{
		lru.ll.MoveToFront(elem) 
		return elem.Value.(*entry).value
	}
	return -1
}

func (lru *LRUCache) Put(key, value int){
	if elem, found := lru.cache[key]; found{
		lru.ll.MoveToFront(elem) 
		elem.Value.(*entry).value = value 
	}else {
		if lru.ll.Len() >= lru.capacity {
			tail := lru.ll.Back()
			if tail != nil {
				delete(lru.cache, tail.Value.(*entry).key)
				lru.ll.Remove(tail)
			}
		}
		e := &entry{key, value}
		ele := lru.ll.PushFront(e)
		lru.cache[key] = ele 
	}
}

func main(){
	lru := NewLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))	
	fmt.Println(lru.Get(4))
}
