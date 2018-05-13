package main

import (
	"log"
	"testing"
)

func TestRequestAPIGW(t *testing.T) {
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

	_, err := RequestAPIGW(params)
	if err != nil {
		t.Error("RequestAPIGW: error")
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
		t.Error("RequestAPIGW: error")
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
	log.Printf("%#v", expectParams)
	log.Printf("%#v", structParams)

	if expectParams != structParams {
		t.Error("Generated struct are not equal")
	}
}
