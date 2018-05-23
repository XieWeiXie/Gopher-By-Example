package download

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorNil       = errors.New("response is  nil")
	ErrorWrongCode = errors.New("http response code is wrong")
)

func Download(url string) (*goquery.Document, error) {

	var (
		resp *http.Response
		err  error
	)

	if resp, err = http.Get(url); err != nil {
		return nil, ErrorNil
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, ErrorWrongCode
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
