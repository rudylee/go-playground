package basic

import "fmt"

func main() {
	signal := make(chan bool)
	go printSomething(signal)
	fmt.Println("Print without go routine")
	<-signal
}

func printSomething(s chan bool) {
	fmt.Println("Something is printed")
	s <- true
}
