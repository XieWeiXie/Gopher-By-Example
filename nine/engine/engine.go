package engine

import (
	"errors"
	"go-example-for-live/nine/download"

	"fmt"

	"go-example-for-live/nine/parse/gocn"

	"go-example-for-live/nine/domain"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorRun = errors.New("engine error")
)

var (
	ContentsNil = domain.Contents{}
)

func Run(request domain.Request) (domain.Contents, error) {

	var (
		doc *goquery.Document
		err error
	)
	if doc, err = download.Download(request.Url); err != nil {
		fmt.Println(err)
		return ContentsNil, ErrorRun
	}

	return gocn.TitleParse(doc), nil

}
