package main

import (
	"log"

	"strings"

	"bytes"
	"net/http"

	"fmt"

	"encoding/json"

	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type SlackConfig struct {
	URL       string
	Channel   string
	Username  string
	IconEmoji string
}

type SlackJSON struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Text      string `json:"text"`
}

func FetchRSS(url string) (imageList []string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	//htmlないからimgタグを取得
	//imageList = append(imageList, "hoge")
	selecter := doc.Find("item description")
	//log.Println(len(selecter.Nodes))
	for i := range selecter.Nodes {
		node := selecter.Eq(i)
		text := node.Text()
		list := fetchDescription(text)
		imageList = append(imageList, list...)
	}
	//log.Println(imageList)

	return imageList
}

//slackへの通知

func PostSlack(config SlackConfig, body string) (err error) {

	//送信するJSONを作成
	sendData := SlackJSON{
		Channel:   config.Channel,
		Username:  config.Username,
		IconEmoji: config.IconEmoji,
		Text:      body,
	}
	jsonBytes, err := CreateJSON(sendData)
	if err != nil {
		return
	}

	//送信処理
	err = PostJSON(jsonBytes, config.URL)
	return
}

func PostData() {
	data := url.Values{}
	data.Set("payload", "{hoge : \"hogehoge\"}")
}

func PostJSON(jsonBytes []byte, siteURL string) (err error) {
	log.Print(siteURL)
	log.Print(string(jsonBytes))

	req, err := http.NewRequest(
		"POST",
		siteURL,
		bytes.NewBuffer(jsonBytes),
		//		bytes.NewBuffer([]byte(string(jsonBytes))),
	)

	if err != nil {
		return
	}
	// Content-Type 設定
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	//log.Println(resp)
	//log.Println(resp.Body)
	return
}

func CreateJSON(sendData SlackJSON) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(sendData)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	//fmt.Println(string(jsonBytes))
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
	//log.Println("img: ", len(selecter.Nodes))
	for i := range selecter.Nodes {
		node := selecter.Eq(i)
		text, _ := node.Attr("src")
		imageList = append(imageList, text)
	}
	return
}
