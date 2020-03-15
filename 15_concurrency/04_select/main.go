package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	/*
		This will print the output like following:

		Every 500ms
		Every 2s
		Every 500ms
		Every 2s
		Every 500ms
		Every 2s

		This is because, even though channel c1 is ready to send message much
		sooner, the main goroutine needs to wait until channel c2 is ready to
		send a message. So, this always results in one channel being executed
		after another.
	*/
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	/*
		But, we may not want this, we may want to recieve
		messages as soon as a channel is ready to recieve it. This is where
		we use the select statement to select whichever channel is ready to
		be recieved from and recieve it.
	*/
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
