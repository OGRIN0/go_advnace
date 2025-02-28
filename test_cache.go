package main 

import(
    "fmt"
    "testing"
    "sync"
)

func TestCache(t *testing.T) {
    cache := NewCache()
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            cache.Set(fmt.Sprint(val), val)
        }(i)
    }

    wg.Wait()

    for i := 0; i < 10; i++ {
        val, exists := cache.Get(fmt.Sprint(i))
        if !exists {
            t.Errorf("Expected value %d to exist", i)
        }
        if val != i {
            t.Errorf("Expected %d, got %v", i, val)
        }
    }
}