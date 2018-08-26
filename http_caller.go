package main

import (
	"bytes"
	"net/http"
)

type HttpCallInterface interface {
	PostJson(jsonBytes []byte, siteURL string) (err error)
}

type HttpCaller struct {
}

func (c *HttpCaller) PostJson(jsonBytes []byte, siteURL string) (err error) {
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
