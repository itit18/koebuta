package main

import (
	"log"
	"testing"
)

func TestConvertResponse(t *testing.T) {
	result, err := ConvertResponse("test")
	if err != nil {
		t.Error("ConvertResponse: error")
	}
	if result.Text != "test" {
		t.Error("text parameter is unknown　value")
		log.Printf("%#v", result.Text)
	}
}

func TestConvertRequest(t *testing.T) {
	params := map[string]string{
		"team_domain":  "slack.domain",
		"user_id":      "U11111111",
		"channel_name": "TestChannel",
		"service_id":   "111111111111",
		"text":         "test message",
		"timestamp":    "1526192788.000000",
		"token":        "ABC111111111111111111111",
		"trigger_word": "word",
		"user_name":    "jon doe",
		"channel_id":   "C11111111",
		"team_id":      "T11111111",
	}

	structParams, err := ConvertRequest(params)
	if err != nil {
		t.Error("ConvertRequest: error")
	}

	expectParams := slackRequest{
		Text:        "test message",
		UserID:      "U11111111",
		UserName:    "jon doe",
		ChannelID:   "C11111111",
		ChannelName: "TestChannel",
		ServiceID:   "111111111111",
		TeamDomain:  "slack.domain",
		TeamID:      "T11111111",
		Timestamp:   "1526192788.000000",
		Token:       "ABC111111111111111111111",
		TriggerWord: "word",
	}

	if expectParams != structParams {
		t.Error("Generated struct are not equal")
		log.Printf("%#v", expectParams)
		log.Printf("%#v", structParams)
	}
}

func TestPostJSON(t *testing.T) {
	sendData := incomingJSON{
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
