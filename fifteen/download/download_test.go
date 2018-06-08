package download

import (
	"fmt"
	"testing"
)

func TestGetHtmlResponse(test *testing.T) {
	tt := []struct {
		url string
	}{
		{
			url: "https://www.v2ex.com/?tab=jobs",
		},
		{
			url: "https://www.jianshu.com/u/58f0817209aa",
		},
	}

	for _, t := range tt {
		doc, err := GetHtmlResponse(t.url)
		if err != nil {
			test.Error("get html response failed")
		} else {
			fmt.Println(doc.Html())
		}

	}
}
