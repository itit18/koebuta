package main

import (
	"testing"
)

func TestPostDaRta(t *testing.T) {
	PostData()
}

func TestFetchRSS(t *testing.T) {
	html := FetchRSS("http://himanji.tumblr.com/rss")
	expectValue := html[0]
	if expectValue == "" {
		t.Error("返り値が空")
	}

	if expectValue[0:4] != "http" {
		t.Error("返り値がURL出ない可能性がある")
	}

}

func TestCreateJSON(t *testing.T) {
	sendData := SlackJSON{
		Channel:   "test1",
		Username:  "test2",
		IconEmoji: "test3",
		Text:      "test4",
	}
	_, err := CreateJSON(sendData)
	if err != nil {
		t.Error("JSONの生成に失敗")
	}
}

func TestPostJSON(t *testing.T) {
	sendData := SlackJSON{
		Channel:   "test1",
		Username:  "test2",
		IconEmoji: "test3",
		Text:      "test4",
	}
	val, err := CreateJSON(sendData)

	err = PostJSON(val, "https://example.com/")
	if err != nil {
		t.Error("api clientが失敗しました")
	}
}
