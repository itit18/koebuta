package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type ImageUrl struct {
	list          []string
	externalSites []string
}

//TODO: 動作テストする / もしくはちゃんとテスト書く

func (iu *ImageUrl) SetExternalSites(url []string) {
	iu.externalSites = nil
	if len(url) == 0 {
		iu.externalSites = append(iu.externalSites, "http://himanji.tumblr.com/rss")
		iu.externalSites = append(iu.externalSites, "http://pocapontas.tumblr.com/rss")
		iu.externalSites = append(iu.externalSites, "https://hiyayall.tumblr.com/rss")
		iu.externalSites = append(iu.externalSites, "http://maeda-toshiie.tumblr.com/rss")
		iu.externalSites = append(iu.externalSites, "http://ktminamotokr.tumblr.com/rss")
	} else {
		for _, v := range url {
			iu.externalSites = append(iu.externalSites, v)
		}
	}
}

func (iu *ImageUrl) FetchImageFromExternal() (err error) {
	//TODO: バッファなしchannelの実装がよくわからないので大きいバッファを指定している
	imagesChan := make(chan string, 10000)
	errorChan := make(chan error, 10000)
	wg := &sync.WaitGroup{}

	// 並列にfetchを実行
	for _, v := range iu.externalSites {
		wg.Add(1)
		log.Println(v)
		go iu.fetch(v, imagesChan, errorChan, wg)
	}

	wg.Wait()
	close(errorChan)
	close(imagesChan)

	//ここのループはchannelが持っているデータ分だけ
	for err = range errorChan {
		if err != nil {
			return
		}
	}
	for image := range imagesChan {
		iu.list = append(iu.list, image)
	}

	return
}

func (iu *ImageUrl) Len() int {
	return len(iu.list)
}

func (iu *ImageUrl) GetRandom() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(iu.list))
	return iu.list[i]
}

func (iu *ImageUrl) fetch(url string, c chan string, errChannel chan error, wg *sync.WaitGroup) {
	list, err := FetchRSS(url)
	if err != nil {
		errChannel <- err
	}
	for _, v := range list {
		c <- v
	}

	defer wg.Done()
}
