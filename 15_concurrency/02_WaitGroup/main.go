package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroup is basically just a special counter. It keeps count
	// of the number of goroutines launched.
	var wg sync.WaitGroup

	// Increase the count by 1 before launching the goroutine.
	// Increasing the count inside the goroutine might not work at all,
	// because the "main" goroutine executes the next line after launching
	// the new goroutine almost immediately, so in our case, calling wg.Wait()
	// will not have any effect since the counter did not increase at all
	// i.e the newly launched goroutine haven't even started its execution
	wg.Add(1)

	go func() {
		// The execution inside this goroutine is synchronous, so count()
		// will execute synchronously
		count("sheep")

		// Done with this goroutine, decrement the counter
		wg.Done()
	}()

	// This will make the current goroutine wait until WaitGroup's counter
	// goes back to 0 i.e all goroutines have finished their execution.
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
