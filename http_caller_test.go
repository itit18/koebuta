package main

import (
	"encoding/json"
	"testing"
)

func TestHttpCaller_PostJson(t *testing.T) {
	sendData := map[string]string{
		"Key":  "hoge",
		"Key2": "fuga",
	}

	val, err := json.Marshal(sendData)
	if err != nil {
		t.Error("JSON Marshal error:", err)
	}

	client := &HttpCaller{}
	err = client.PostJson(val, "https://example.com/")
	if err != nil {
		t.Error("外部URLへのhttpリクエストが失敗しました")
	}
}
