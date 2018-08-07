package main

import (
	"fmt"
	"time"
)

func main() {
	signal := make(chan int)
	go periodicSend(signal)
	for i := range signal {
		fmt.Println(i)
	}
}

func periodicSend(signal chan int) {
	for {
		signal <- 3
		time.Sleep(3 * time.Second)
	}
}
