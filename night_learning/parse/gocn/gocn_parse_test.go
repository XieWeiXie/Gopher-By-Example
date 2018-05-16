package gocn

import (
	"testing"

	"go-example-for-live/night_learning/download"

	"github.com/PuerkitoBio/goquery"
)

func TestTitleParse(t *testing.T) {

	tests := []struct {
		url      string
		document *goquery.Document
	}{
		{
			url: "https://gocn.io/sort_type-new__day-0__is_recommend-0",
		},
		{
			url: "https://gocn.io/sort_type-new__day-0__is_recommend-0__page-3",
		},
	}

	for _, test := range tests {
		test.document, _ = download.Download(test.url)
		TitleParse(test.document)
	}
}

func TestPassageParse(t *testing.T) {

	tests := []struct {
		url string
	}{
		{
			url: "https://gocn.io/article/783",
		},
		{
			url: "https://gocn.io/question/1803",
		},
		{
			url: "https://gocn.io/question/1917",
		},
		{
			url: "https://gocn.io/article/788",
		},
	}

	for _, test := range tests {
		doc, _ := download.Download(test.url)
		PassageParse(doc)
	}

}
