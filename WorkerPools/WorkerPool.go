package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

// Worker function which processes jobs
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		// Simulate work
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
	}
}

func main() {
	const (
		numJobs    = 5
		numWorkers = 3
	)

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs to the job queue
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs processed")
}
