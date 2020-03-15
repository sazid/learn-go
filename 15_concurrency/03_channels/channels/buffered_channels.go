package channels

import "fmt"

// ChannelLockDemo : Demonstrates channel locking
func ChannelLockDemo() {
	// This will try to send a message through the channel but
	// will wait until the reciever is ready to recieve the message.
	// In this case, it'll keep waiting and won't proceed to the next
	// line until the channel message is recieved.
	c := make(chan string)
	c <- "hello"

	msg := <-c
	fmt.Println(msg)
}

// BufferedChannelDemo : Demonstrates the buffered channel
func BufferedChannelDemo() {
	// This won't block the reciever goroutine until the channel
	// buffer is full, in this case until 2 messages are sent
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
