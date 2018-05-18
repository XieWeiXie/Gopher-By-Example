package engine

import (
	"errors"
	"go-example-for-live/nine_learning/download"

	"fmt"

	"go-example-for-live/nine_learning/parse/gocn"

	"go-example-for-live/nine_learning/domain"

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
