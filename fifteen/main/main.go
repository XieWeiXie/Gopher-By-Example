package main

import (
	"go-example-for-live/fifteen/engine"
	"go-example-for-live/fifteen/parse/jianshu"
	"go-example-for-live/fifteen/parse/v2ex"
)

func main() {
	engine.Run(
		engine.Request{
			Url:       "https://www.v2ex.com/?tab=jobs",
			ParseFunc: v2ex.JobParse,
		},
		engine.Request{
			Url:       "https://www.jianshu.com/u/58f0817209aa",
			ParseFunc: jianshu.InfoParse,
		},
	)
}
