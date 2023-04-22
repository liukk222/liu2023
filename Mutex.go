package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	counter int
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter--
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go decrement(&wg)
	}
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
