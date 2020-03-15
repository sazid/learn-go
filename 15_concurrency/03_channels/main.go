package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sazid/go_crash_course/15_concurrency/03_channels/channels"
)

func main() {
	/*
		A channel is a way for goroutines to communicate with each other
	*/

	// Create a channel
	c := make(chan string)
	go count("sheep", c)

	/*
		Note: Sending messages through the channels is a synchronous operation,
		i.e if the sender reaches the line where it needs to send a message,
		it'll wait until the reciever's execution reaches the line where its
		ready to recieve a message from the channel.

		This synchronous message passing mechanism is what enables a very easy
		and straight forward way to communicate between the different
		goroutines.
	*/

	/*
		// See the alternative below once you read through this one first.

		for {
			// Recieve a message from the channel
			// msg := <-c

			// The channel will also send a second value which indicates whether
			// the channel is still open or not.
			msg, open := <-c

			// We're done recieving messages
			if !open {
				break
			}

			fmt.Println(msg)
		}
	*/

	// Alternative way - syntactic sugar
	// Iterate over the range of messages recieved from the channel
	// Note that the loop will automatically break once the channel is closed.
	for msg := range c {
		fmt.Println(msg)
	}

	// channels.ChannelLockDemo()
	channels.BufferedChannelDemo()
}

/*
	Channels are like pipes which you can use to send messages back and forth.
	They're declared with the following syntax:
	name chan type
	where "type" denotes what type of channel it is going to be. In this case,
	c is a string channel and it can only send and recieve string messages.
	However, any "type" is permitted, you can even use channels of type
	channels to pass channels as messages.
*/
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		msg := thing + " " + strconv.Itoa(i)

		// Pass a msg into the channel
		c <- msg

		time.Sleep(time.Millisecond * 500)
	}

	/*
		We need to close the channel because once the execution of this
		function is finished, we expect that the reciever goroutine should not
		wait for any new messages from the channel.

		REMEMBER: The reciever must never close() the channel because on the
		reciever end, we do not know whether the channel is finished sending
		all of its messages. Calling close() in reciever will prematurely
		close the channel however the sender may still try to send messages
		which will result in a panic.
	*/
	close(c)
}
