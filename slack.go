package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type SlackRequest struct {
	Text        string
	UserID      string
	UserName    string
	ChannelID   string
	ChannelName string
	ServiceID   string
	TeamDomain  string
	TeamID      string
	Timestamp   string
	Token       string
	TriggerWord string
}

//lambdaからAPI GWに返り値を渡す際にjsonフォーマットに変換されるので、key名を指定してる
type SlackResponse struct {
	Text string `json:"text"`
}

type SlackConfig struct {
	URL       string
	Channel   string
	Username  string
	IconEmoji string
}

type incomingJSON struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Text      string `json:"text"`
}

// outgoing web hook

//slackに対応したresponse形式に変換
func ConvertResponse(msg string) (res SlackResponse, err error) {
	res = SlackResponse{
		Text: msg,
	}

	return
}

// AWS GWから渡ってくる値はmap型なのでstructに変換する
func ConvertRequest(params map[string]string) (res SlackRequest, err error) {
	res = SlackRequest{
		Text:        params["text"],
		UserID:      params["user_id"],
		UserName:    params["user_name"],
		ChannelID:   params["channel_id"],
		ChannelName: params["channel_name"],
		ServiceID:   params["service_id"],
		TeamDomain:  params["team_domain"],
		TeamID:      params["team_id"],
		Timestamp:   params["timestamp"],
		Token:       params["token"],
		TriggerWord: params["trigger_word"],
	}

	return
}

func Authentication(token string) error {
	if token != os.Getenv("KB_SLACK_TOKEN") {
		return errors.New("token do not match")
	}

	return nil
}

// incoming web hook

func PostSlack(config SlackConfig, body string) (err error) {

	//送信するJSONを作成
	sendData := incomingJSON{
		Channel:   config.Channel,
		Username:  config.Username,
		IconEmoji: config.IconEmoji,
		Text:      body,
	}

	jsonBytes, err := json.Marshal(sendData)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	//送信処理
	err = PostJSON(jsonBytes, config.URL)

	return
}

func PostJSON(jsonBytes []byte, siteURL string) (err error) {
	req, err := http.NewRequest(
		"POST",
		siteURL,
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		return
	}
	// Content-Type 設定
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return
}

func CreateIncomingConfig() (config SlackConfig) {
	config = SlackConfig{
		URL:       os.Getenv("KB_URL"),
		Username:  os.Getenv("KB_USER"),
		IconEmoji: os.Getenv("KB_ICON"),
		Channel:   os.Getenv("KB_CHANNEL"),
	}

	return
}
