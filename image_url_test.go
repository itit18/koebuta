package main

import (
	"log"
	"testing"
)

//RSS取得処理のモッククライアントを定義
type MockRssClient struct{}

func (c *MockRssClient) Fetch(url string) (imageList []string, err error) {
	imageList = append(imageList, "hoge")
	return
}

func TestImageUrl_FetchImageFromExternal(t *testing.T) {
	iu := ImageUrl{}
	iu.client = &MockRssClient{}
	url := []string{
		"http://maeda-toshiie.tumblr.com/rss",
		"http://ktminamotokr.tumblr.com/rss",
	}
	iu.SetExternalSites(url)
	err := iu.FetchImageFromExternal()
	if err != nil {
		t.Error("FetchImageFromExternal: error")
		log.Print(err)
	}
	if 0 == iu.Len() {
		t.Error("do not fetch image url")
	}
}
