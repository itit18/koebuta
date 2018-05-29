package main

import (
	"context"
	"log"

	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(koebuta)
}

//runner
func koebuta(ctx context.Context, params map[string]string) (res slackResponse, err error) {
	log.Print(ctx)
	log.Print(params)

	res, err = outGoingHook(params)

	return
}

func outGoingHook(params map[string]string) (res slackResponse, err error) {
	structParams, err := ConvertRequest(params)
	if err != nil {
		return
	}
	log.Printf("%#v", structParams)

	image, err := FetchImageURL()
	if err != nil {
		return
	}

	res, err = ConvertResponse(image)
	if err != nil {
		return
	}

	return
}

func inComingHook() (string, error) {
	image, err := FetchImageURL()
	if err != nil {
		return "error", err
	}

	config := CreateIncomingConfig()
	err = PostSlack(config, image)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("success"), nil
}
