package main

import (
	"github.com/charlsonz/crawler/config"
	"github.com/charlsonz/crawler/engine"
	"github.com/charlsonz/crawler/persist"
	"github.com/charlsonz/crawler/scheduler"
	"github.com/charlsonz/crawler/xcar/parser"
)

const urlCar = "https://www.xcar.com.cn/"

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    urlCar,
		Parser: engine.NewFuncParser(parser.ParseCarList, config.ParseCarList),
	})
}
