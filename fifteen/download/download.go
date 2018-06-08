package download

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrRequest  = errors.New("request error")
	ErrResponse = errors.New("response error")
)

func GetHtmlResponse(url string) (*goquery.Document, error) {
	var (
		request *http.Request
		err     error
	)
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, ErrRequest
	}

	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	var client *http.Client
	client = http.DefaultClient

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return nil, ErrResponse
	}

	defer response.Body.Close()
	return goquery.NewDocumentFromReader(response.Body)
}
