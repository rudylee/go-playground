package main

import (
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8000", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
}
