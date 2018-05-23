package main

import (
	"go-example-for-live/seven/engine"
	"go-example-for-live/seven/parse/github"
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
