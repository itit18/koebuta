package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(koebuta)
}

func koebuta() {
	// 複数のリソースURLに対応する
	sites := [5]string{
		"http://himanji.tumblr.com/rss",
		"http://pocapontas.tumblr.com/rss",
		"https://hiyayall.tumblr.com/",
		"http://maeda-toshiie.tumblr.com/rss",
		"http://ktminamotokr.tumblr.com/rss",
	}
	images := []string{} //TODO: sliceの大きさを指定するとエラーになるのはなぜ…
	for _, v := range sites {
		list := FetchRSS(v)
		images = append(images, list...)
	}
	log.Println(len(images))

	config := SlackConfig{
		URL:       os.Getenv("KB_URL"),
		Username:  os.Getenv("KB_USER"),
		IconEmoji: os.Getenv("KB_ICON"),
		Channel:   os.Getenv("KB_CHANNEL"),
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(images))
	err := PostSlack(config, images[i])
	if err != nil {
		log.Fatal(err)
	}
}
