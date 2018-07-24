package main

import (
	"log"
	"testing"
)

//TODO: 外部API依存の処理をなくす
func TestImageUrl_FetchImageFromExternal(t *testing.T) {
	iu := ImageUrl{}
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
