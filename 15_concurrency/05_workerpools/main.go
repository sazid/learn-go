package main

import "fmt"

func main() {
	/*
		Worker Pool pattern:
		A queue of work to be done is put inside a pool and then each of the
		tasks are taken from the pool and executed.
	*/

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Launch multiple worker goroutines which will take jobs from the job queue
	// and perform the jobs/tasks independently
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		// Send jobs to the channel
		jobs <- i
	}
	close(jobs)

	// Recieve the results
	for msg := range results {
		fmt.Println(msg)
	}
}

// jobs : jobs to do (recieve only channel)
// results : send results from the jobs (send only channel)
func worker(jobs <-chan int, results chan<- int) {
	// Take jobs from the job channel
	for n := range jobs {
		// Perform the currently selected job and send the result of the
		// job through the "results" channel.
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
