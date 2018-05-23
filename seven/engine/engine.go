package engine

import (
	"errors"
	"fmt"
	"go-example-for-live/seven/download"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorDocWrong = errors.New("document wrong")
)

type Trending struct {
}

func (t Trending) Run(request RequestForGithub) {

	var doc *goquery.Document

	doc, err := download.Download(request.URL)
	if err != nil {
		fmt.Println(ErrorDocWrong)
		return
	}
	if doc != nil {
		fmt.Println("Game start!")
		request.ParseFunc(doc)
	} else {
		fmt.Println("Game over!")
	}

}
