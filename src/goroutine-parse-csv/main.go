package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	tasks := make(chan string)

	csvFile, err := os.Open("education.csv")
	if err != nil {
		log.Fatal(err)
	}

	go worker(tasks, 1)
	go worker(tasks, 2)

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		// Skip headers
		if line[0] == "ServiceApprovalNumber" {
			continue
		}

		tasks <- line[2]
	}
	close(tasks)
}

func worker(tasks <-chan string, id int) {
	for n := range tasks {
		fmt.Println("Worker", id, "working on", n)
	}
}
