package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type MyData struct {
	PublishingOffice string
	ReportDatetime string
	TargetArea string
	HeadlineText string
	Text string
}

func main() {
	p := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	var item MyData
	er = json.Unmarshal(s, &item)
	if er != nil {
		panic(er)
	}

	println("エリア：", item.TargetArea, "\n", item.Text)
}
