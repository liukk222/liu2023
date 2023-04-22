/*
	package main

import (

	"fmt"
	"sync"

)

var (

	count int
	m     sync.Mutex // 互斥锁

)

	func increment(wg *sync.WaitGroup) {
		for i := 0; i < 1000; i++ {
			m.Lock() // 加锁
			count++
			m.Unlock() // 解锁
		}
		wg.Done()
	}

	func main() {
		var wg sync.WaitGroup
		wg.Add(2)

		go increment(&wg)
		go increment(&wg)

		wg.Wait()

		fmt.Println(count)
	}
*/
