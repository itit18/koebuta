package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchRSS(url string) (imageList []string, err error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	//html内からimgタグを取得
	selecter := doc.Find("item description")
	for i := range selecter.Nodes {
		node := selecter.Eq(i)
		text := node.Text()
		list := fetchDescription(text)
		imageList = append(imageList, list...)
	}

	return
}

func CreateJSON(sendData incomingJSON) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(sendData)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	return
}

//private

func fetchDescription(text string) (imageList []string) {
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
