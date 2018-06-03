package main

import (
	"math/rand"
	"time"
)

func FetchImageURL() (imageURL string, err error) {
	sites := [5]string{
		"http://himanji.tumblr.com/rss",
		"http://pocapontas.tumblr.com/rss",
		"https://hiyayall.tumblr.com/",
		"http://maeda-toshiie.tumblr.com/rss",
		"http://ktminamotokr.tumblr.com/rss",
	}
	images := []string{} //TODO: sliceの大きさを指定するとエラーになるのはなぜ…
	var list []string
	for _, v := range sites {
		list, err = FetchRSS(v)
		if err != nil {
			return
		}
		images = append(images, list...)
	}
	//log.Println(len(images))

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(images))
	imageURL = images[i]

	return
}
