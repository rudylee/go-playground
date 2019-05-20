package main

import "fmt"

// func main() {
// 	c := make(chan string)
//
// 	go count("test", c)
// 	for msg := range c {
// 		fmt.Println(msg)
// 	}
// }
//
// func count(name string, c chan string) {
// 	for i := 1; i < 5; i++ {
// 		c <- name
// 		time.Sleep(time.Millisecond * 300)
// 	}
//
// 	close(c)
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
//
// 	go func() {
// 		for {
// 			c1 <- "Every 500ms"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()
//
// 	go func() {
// 		for {
// 			c2 <- "Every 2000ms"
// 			time.Sleep(time.Millisecond * 2000)
// 		}
// 	}()
//
// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
