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
	mode := os.Getenv("KB_MODE")
	ary := [4]string{"outgoing", "incoming", "stock", "simple"}
	for _, v := range ary {
		if v == mode {
			k.mode = mode
			break
		}
	}

	if len(k.mode) == 0 {
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
	switch k.mode {
	case "outgoing":
		err = k.outGoingHook(k.lambdaParams)
	case "incoming":
		err = k.inComingHook()
	case "stock":
		err = k.stock()
	case "simple":
		err = k.simple()
	default:
		err = errors.New("mode unknown: execute `SetMode`")
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

	iu := &ImageUrl{}
	iu.SetExternalSites([]string{})
	err = iu.FetchImageFromExternal()
	if err != nil {
		return
	}
	image := iu.GetRandom()

	res, err := ConvertResponse(image)
	if err != nil {
		return
	}

	k.responseJSON = res
	return nil
}

func (k *Koebuta) inComingHook() error {
	iu := &ImageUrl{}
	iu.SetExternalSites([]string{})
	err := iu.FetchImageFromExternal()
	if err != nil {
		return err
	}
	image := iu.GetRandom()

	config := CreateIncomingConfig()
	err = PostSlack(config, image)
	if err != nil {
		log.Fatal(err)
	}

	k.responseText = "success"
	return nil
}

func (k *Koebuta) stock() error {
	iu := &ImageUrl{}
	iu.SetExternalSites([]string{})
	err := iu.FetchImageFromExternal()
	if err != nil {
		return err
	}
	image := iu.GetRandom()
	log.Println(image)
	k.responseText = "success"

	return nil
}

func (k *Koebuta) simple() error {
	iu := &ImageUrl{}
	iu.SetExternalSites([]string{})
	err := iu.FetchImageFromExternal()
	if err != nil {
		return err
	}
	image := iu.GetRandom()
	k.responseText = image

	return nil
}
