package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	defer os.Remove(tmpFile.Name())
}
