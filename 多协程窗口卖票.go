/* package main

import (
	"fmt"
	"sync"
)

func sellTickets(wg *sync.WaitGroup, ch chan int, id int) {
	defer wg.Done()
	for {
		ticket, ok := <-ch
		if !ok {
			fmt.Printf("goroutine %d: channel is closed\n", id)
			return
		}
		fmt.Printf("goroutine %d: sell ticket %d\n", id, ticket)
	}
}

func main() {
	const numTickets = 100
	const numSellers = 4

	var wg sync.WaitGroup
	wg.Add(numSellers)

	ch := make(chan int, numTickets)
	for i := 1; i <= numTickets; i++ {
		ch <- i
	}
	close(ch)

	for i := 1; i <= numSellers; i++ {
		go sellTickets(&wg, ch, i)
	}

	wg.Wait()
}
*/