package main

import (
	"fmt"
	"time"
)

/*
	Concurrency is not the same as parallelism.
	To run things in parallel means to run two things at the
	same time (i.e two things executing in two different physical cores).

	Concurrent tasks *can* run at the same time, but is not guaranteed.

	Go provides tools to make it easier to write concurrent programs.
	Then, it decides itself whether to run tasks concurrently or not
	depending on the cpu its running on, that is, we don't havae to
	deal with managing the concurrent tasks ourselves.
*/

func main() {
	/*
		If we call the function with the keyword "go" in front of it,
		it won't wait for the function to finish before moving on to
		the next line i.e it launches/runs the function in the
		background and continues on the next line immediately.

		"go" keyword creates a goroutine and this goroutine runs
		concurrently. You can launch as many goroutines as you like:
		hundreds or thousands of them.
	*/

	// There are now two goroutines - the main go routine and our
	// own newly launched goroutine.
	go count("sheep")

	// Uncommenting this line and commenting the next one will
	// actually quit the program as soon as it starts. Why though?
	// Its because, once the execution in the main goroutine is
	// finished, the program quits, it doesn't wait for other goroutines
	// to finish their execution.

	// go count("fish")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
