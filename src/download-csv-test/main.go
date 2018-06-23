package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	DownloadCSV("https://www.asx.com.au/asx/research/ASXListedCompanies.csv", "asx-companies.csv")
}

func DownloadCSV(url string, filename string) error {
	out, _ := os.Create(filename)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Fatal(err)
	}

	return err
}
