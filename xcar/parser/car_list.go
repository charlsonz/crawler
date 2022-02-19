package parser

import (
	"regexp"

	"github.com/charlsonz/crawler/config"
	"github.com/charlsonz/crawler/engine"
)

const host = "http://newcar.xcar.com.cn"

var (
	carModelRE = regexp.MustCompile(`<a href="(/\d+/)" target="_blank" class="list_img">`)
	carListRE  = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)"`)
)

func ParseCarList(contents []byte, _ string) engine.ParseResult {
	matches := carModelRE.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    host + string(m[1]),
			Parser: engine.NewFuncParser(ParseCarModel, config.ParseCarModel),
		})
	}

	matches = carListRE.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    "http:" + string(m[1]),
			Parser: engine.NewFuncParser(ParseCarList, config.ParseCarList),
		})
	}

	return result
}
