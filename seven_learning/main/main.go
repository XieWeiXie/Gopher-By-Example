package main

import (
	"go-example-for-live/seven_learning/engine"
	"go-example-for-live/seven_learning/parse/github"
)

func main() {

	var simplerTest engine.Trending

	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending",
			ParseFunc: github.ParseForGithub,
		},
	)
	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending/developers",
			ParseFunc: github.ParseForDevelopers,
		},
	)

}
