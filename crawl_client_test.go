package main

import (
	"testing"
)

func TestFetchRSS(t *testing.T) {
	client := CrawlClient{}
	html, err := client.Fetch("http://himanji.tumblr.com/rss")
	if err != nil {
		t.Error("FetchRSS: error")
	}
	expectValue := html[0]
	if expectValue == "" {
		t.Error("返り値が空")
	}
	if expectValue[0:4] != "http" {
		t.Error("返り値がURL出ない可能性がある")
	}

}
