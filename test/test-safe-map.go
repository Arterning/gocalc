package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a new sync.Map
	var m sync.Map

	// Add key-value pairs concurrently
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int, value string) {
			m.Store(key, value)
			wg.Done()
		}(i, fmt.Sprintf("value%d", i))
	}
	wg.Wait()

	// Read from the map concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			if value, ok := m.Load(key); ok {
				fmt.Printf("Value for key %d: %s\n", key, value)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}