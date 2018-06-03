package main

import (
	"log"
	"testing"
)

func TestFetchImageURL(t *testing.T) {
	url, err := FetchImageURL()
	if err != nil {
		t.Error("FetchImageURL: error")
		log.Print(err)
	}
	log.Printf(url)
}
