package engine

import (
	"go-example-for-live/nine/download"

	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Request struct {
	Url       string
	ParseFunc func(*goquery.Document) []string
}

type Requests []Request

var (
	ErrDownload = errors.New("download err")
)

func Run(seeds ...Request) {
	var requests Requests
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := download.Download(r.Url)
		if err != nil {
			fmt.Println(ErrDownload)
			continue
		}
		fmt.Println(r.ParseFunc(body))

	}

}
