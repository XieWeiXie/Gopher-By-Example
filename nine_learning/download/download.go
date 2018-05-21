package download

import (
	"net/http"

	"errors"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrorHttpRequest  = errors.New("http request error")
	ErrorHttpResponse = errors.New("http response error")
)

func Download(url string) (*goquery.Document, error) {
	var (
		req *http.Request
		err error
	)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, ErrorHttpRequest
	}

	client := http.DefaultClient
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	var (
		resp *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return nil, ErrorHttpResponse
	}

	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

func Downloader(url string) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, ErrorHttpRequest
	}

	client := http.DefaultClient
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	var (
		resp *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return nil, ErrorHttpResponse
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
