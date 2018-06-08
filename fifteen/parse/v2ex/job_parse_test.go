package v2ex

import (
	"testing"

	"go-example-for-live/nine/download"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func TestV2exParse(test *testing.T) {
	docV2ex, _ := download.Download("https://www.v2ex.com/?tab=jobs")

	tt := []struct {
		doc *goquery.Document
	}{
		{
			doc: docV2ex,
		},
	}

	for _, t := range tt {
		fmt.Println(JobParse(t.doc))
	}
}
