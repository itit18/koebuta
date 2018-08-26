package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CrawlClient struct {
}

func (c *CrawlClient) Fetch(url string) (imageList []string, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	//html内からimgタグを取得
	selecter := doc.Find("item description")
	for i := range selecter.Nodes {
		node := selecter.Eq(i)
		text := node.Text()
		list := c.fetchDescription(text)
		imageList = append(imageList, list...)
	}

	return
}

//private

func (c *CrawlClient) fetchDescription(text string) (imageList []string) {
	r := strings.NewReader(text)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	selecter := doc.Find("img")
	for i := range selecter.Nodes {
		node := selecter.Eq(i)
		text, _ := node.Attr("src")
		imageList = append(imageList, text)
	}

	return
}
