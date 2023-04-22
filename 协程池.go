package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID   int
	Data interface{}
}

func worker(id int, tasks <-chan Task, results chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		results <- task.Data // 模拟任务处理结果
	}
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 4
	tasks := make(chan Task, 10)
	results := make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		tasks <- Task{ID: i, Data: fmt.Sprintf("Task %d", i)}
	}
	close(tasks)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		result := <-results
		fmt.Printf("Result %d: %v\n", i, result)
	}
}
