package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	run()
}

func run() {
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
		URL:       "https://hooks.slack.com/services/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",
		Username:  "test",
		IconEmoji: ":ghost:",
		Channel:   "#apitest",
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(images))
	err := PostSlack(config, images[i])
	if err != nil {
		log.Fatal(err)
	}
}
