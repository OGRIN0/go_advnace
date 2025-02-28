package main

import (
	"log"
	"sync"
	"hash/fnv"
	"fmt"
)

type Shard struct {
	sync.RWMutex 
	data map[string]any 
}

type ShardMap []*Shard 

type Cache struct {
	sync.RWMutex
	data map[string]any
}

func NewShardMap(n int) ShardMap {
	shards := make(ShardMap, n)
	for i := 0; i < n; i++ {
		shards[i] = &Shard{
			data: make(map[string]any),
		}
	}
	return shards
}

func (sm ShardMap) getShard(key string) *Shard {
    hash := fnv.New32()
    hash.Write([]byte(key))
    return sm[hash.Sum32()%uint32(len(sm))]
}

func (sm ShardMap) Get(key string) (any, bool) {
    shard := sm.getShard(key)
    shard.RLock()
    defer shard.RUnlock()
    
    val, exists := shard.data[key]
    return val, exists
}

func (sm ShardMap) Set(key string, value any) {
    shard := sm.getShard(key)
    shard.Lock()
    defer shard.Unlock()
    
    shard.data[key] = value
}

func (sm ShardMap) Delete(key string) {
    shard := sm.getShard(key)
    shard.Lock()
    defer shard.Unlock()
    
    delete(shard.data, key)
}

func (sm ShardMap) Keys() []string {
    keys := make([]string, 0)
    
    for _, shard := range sm {
        shard.RLock()
        for key := range shard.data {
            keys = append(keys, key)
        }
        shard.RUnlock()
    }
    
    return keys
}

func NewCache() Cache{
	return Cache{
		data: make(map[string]any),
	}
}

func (m *Cache) Get(key string) (any, bool) {
    m.RLock()
    defer m.RUnlock()

    val, exists := m.data[key]
    return val, exists
}

func (m *Cache) Set(key string, val any){
	m.Lock()
	defer m.Unlock()

	m.data[key] = val 
}

func (m *Cache) Delete(key string) {
    m.Lock()
    defer m.Unlock()

    delete(m.data, key)
}

func (m *Cache) Contains(key string) bool {
    m.RLock()
    defer m.RUnlock()
    
    _, exists := m.data[key]
    return exists
}

func (m *Cache) Keys() []string {
    m.RLock()
    defer m.RUnlock()

    keys := make([]string, 0, len(m.data))
    for k := range m.data {
        keys = append(keys, k)
    }
    return keys
}

func RunCacheExample(){
	cache := NewCache()

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	keys := cache.Keys()
	for k := range keys {
		log.Printf("key: %v", k)
	}

	a, _ := cache.Get("a")
	log.Printf("a: %v", a)

	b, _ := cache.Get("b")
	log.Printf("b: %v", b)

	z,_ := cache.Get("z")
	log.Printf("z: %v", z)

	cache.Delete("a")
	cache.Delete("z")

	a, exists := cache.Get("a")
	log.Printf("a: %v, exists: %v", a, exists)

	keys = cache.Keys()
	for k := range keys {
		log.Printf("Key: %v", k)
	}
}

// Example usage function
func RunShardMapExample() {
    shards := NewShardMap(8) // Create with 8 shards
    
    // Test concurrent operations
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            key := fmt.Sprintf("key-%d", val)
            shards.Set(key, val)
        }(i)
    }
    wg.Wait()
    
    // Print some results
    log.Printf("Total keys: %d", len(shards.Keys()))
    val, exists := shards.Get("key-1")
    log.Printf("key-1: %v, exists: %v", val, exists)
}

func main() {
    // Run normal cache example
    fmt.Println("=== Running Regular Cache Example ===")
    RunCacheExample()

    fmt.Println("\n=== Running Sharded Cache Example ===")
    RunShardMapExample()
}