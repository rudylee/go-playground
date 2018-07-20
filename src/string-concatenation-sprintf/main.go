package main

import "fmt"

func main() {
	page := "index.html"
	url := fmt.Sprintf("http://www.test.com/%s", page)

	fmt.Println(url)
}
