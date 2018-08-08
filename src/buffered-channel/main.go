package main

import (
	"fmt"
	"time"
)

func main() {
	buffchannel := make(chan int, 5)
	buffchannel <- 3
	buffchannel <- 2
	fmt.Println(<-buffchannel)
	fmt.Println(<-buffchannel)

	fmt.Println(<-waitAndSend(1, 5))

	ic := make(chan int)
	select {
	case fast := <-waitAndSend(3, 1):
		fmt.Println(fast)
	case slow := <-waitAndSend(5, 3):
		fmt.Println(slow)
	case ic <- 23:
		fmt.Println("ic received a value")
	default:
		fmt.Println("all channels are slow")
	}
}

func waitAndSend(v, delay int) chan int {
	channel := make(chan int)

	go func() {
		time.Sleep(time.Duration(delay) * time.Second)
		channel <- v
	}()

	return channel
}
