package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadCSV(t *testing.T) {
	test_csv, _ := ioutil.ReadFile("testdata/asx-companies.csv")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(test_csv)
	}))
	defer ts.Close()

	file := "test.csv"
	err := DownloadCSV(ts.URL, file)

	if err != nil {
		t.Errorf("Shouldn't have received an error, got %s", err)
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("Should have created a CSV file")
	}

	expected, _ := ioutil.ReadFile(file)

	if !bytes.Equal(test_csv, expected) {
		t.Errorf("CSV file should have correct content")
	}

	os.Remove(file)
}
