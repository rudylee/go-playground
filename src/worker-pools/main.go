package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(1, jobs, results)
	go worker(2, jobs, results)

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
