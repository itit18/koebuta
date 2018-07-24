package main

import (
	"os"
	"testing"
)

func TestKoebuta_SetLambdaParams(t *testing.T) {
	expectMessage := "hogehoge"
	k := Koebuta{}
	p := map[string]string{"key": expectMessage}
	k.SetLambdaParams(p)
	if k.lambdaParams["key"] != expectMessage {
		t.Error("lambdaParamsが設定できていない")
	}
}

func TestKoebuta_SetMode(t *testing.T) {
	mode := "outgoing"
	os.Setenv("KB_MODE", mode)
	k := Koebuta{}
	err := k.SetMode()
	if err != nil {
		t.Error(err)
	}
}

func TestStock(t *testing.T) {
	k := Koebuta{}
	err := k.stock()
	if err != nil {
		t.Error(err)
	}
}
