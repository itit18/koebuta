package main

type RssClient interface {
	Fetch(string) ([]string, error)
}
