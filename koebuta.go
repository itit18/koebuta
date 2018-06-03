//koebutaのメインロジック

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Koebuta struct {
	responseText string
	responseJSON SlackResponse
	mode         string
	lambdaParams map[string]string
}

func (k *Koebuta) SetMode() (err error) {
	switch m := os.Getenv("KB_MODE"); m {
	case "outgoing":
		k.mode = "outgoing"
	case "incoming":
		k.mode = "incoming"
	default:
		err = errors.New("mode unknown: set environment variable `KB_MODE`")
	}
	return
}

func (k *Koebuta) SetLambdaParams(params map[string]string) {
	k.lambdaParams = params
}

func (k *Koebuta) Run() (err error) {
	log.Printf("%s", "run")
	log.Printf("%#v", k.lambdaParams)
	switch m := os.Getenv("KB_MODE"); m {
	case "outgoing":
		err = k.outGoingHook(k.lambdaParams)
	case "incoming":
		err = k.inComingHook()
	default:
		err = errors.New("model unknown: set environment variable `KB_MODE`")
	}
	return
}

func (k *Koebuta) FormatResponseToLambda() (r SlackResponse, err error) {
	if len(k.responseText) != 0 {
		str := fmt.Sprintf("\"text\": \"%s\"", k.responseText)
		b := []byte(str)
		json.Unmarshal(b, &r)
		return
	} else if len(k.responseJSON.Text) != 0 { // TODO: lambda環境でちゃんと使えるかを確認する
		r = k.responseJSON
		return
	}

	return r, errors.New("response unknown")
}

//private

func (k *Koebuta) outGoingHook(params map[string]string) (err error) {
	structParams, err := ConvertRequest(params)
	if err != nil {
		return
	}
	log.Printf("%#v", structParams)

	err = Authentication(structParams.Token)
	if err != nil {
		return
	}

	image, err := FetchImageURL()
	if err != nil {
		return
	}

	res, err := ConvertResponse(image)
	if err != nil {
		return
	}

	k.responseJSON = res
	return nil
}

func (k *Koebuta) inComingHook() error {
	image, err := FetchImageURL()
	if err != nil {
		return err
	}

	config := CreateIncomingConfig()
	err = PostSlack(config, image)
	if err != nil {
		log.Fatal(err)
	}

	k.responseText = "success"
	return nil
}
