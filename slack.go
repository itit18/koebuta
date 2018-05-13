package main

import "log"

type slackRequest struct {
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
type slackResponse struct {
	Text string `json:"text"`
}

//AWS API GWを利用する際に必要な諸々の処理
func RequestAPIGW(params map[string]string) (res slackResponse, err error) {
	//paramsのコンバート
	structParams, err := ConvertRequest(params)
	if err != nil {
		return res, err
	}
	log.Printf("%#v", structParams)
	//tokenのチェック
	//request用の構造体を作成
	res = slackResponse{
		Text: "success!",
	}

	return
}

// AWS GWから渡ってくる値はmap型なのでstructに変換する
func ConvertRequest(params map[string]string) (res slackRequest, err error) {
	res = slackRequest{
		Text:        params["text"],
		UserID:      params["user_id"],
		UserName:    params["user_name"],
		ChannelID:   params["channel_id"],
		ChannelName: params["channel_name"],
		ServiceID:   params["service_id"],
		TeamDomain:  params["text"],
		TeamID:      params["team_id"],
		Timestamp:   params["timestamp"],
		Token:       params["token"],
		TriggerWord: params["trigger_word"],
	}

	return
}
