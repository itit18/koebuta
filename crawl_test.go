package main

import (
	"testing"
)

func TestFetchRSS(t *testing.T) {
	html, err := FetchRSS("http://himanji.tumblr.com/rss")
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

func TestCreateJSON(t *testing.T) {
	sendData := incomingJSON{
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
