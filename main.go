package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(koebuta)
}

// AWS LambdaとのI/Oのみを担当する / メソッド名はLambdaの縛りなので変えないこと
// API GWにjsonを返す場合、stringを返すと余計なサニタイズが走ってしまうので構造体として返している
func koebuta(ctx context.Context, params map[string]string) (res SlackResponse, err error) {
	log.Print(ctx)
	log.Print(params)

	k := Koebuta{}
	k.SetLambdaParams(params)
	err = k.SetMode()
	if err != nil {
		log.Fatal(err)
	}

	err = k.Run()
	if err != nil {
		log.Fatal(err)
	}

	res, err = k.FormatResponseToLambda()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
	return
}
