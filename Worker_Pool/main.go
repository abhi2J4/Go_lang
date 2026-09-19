package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Create 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	// Tell workers there are no more jobs
	close(jobs)

	// Wait for all workers
	wg.Wait()

	fmt.Println("All jobs completed")
}