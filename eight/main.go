package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrorHttpNil = errors.New("http error")
	ErrorByteNil = errors.New("byte error")
)

type HttpClient interface {
	Get(string) ([]byte, error)
	Post(string, map[string]interface{}) ([]byte, error)
	Put(string, map[string]interface{}) ([]byte, error)
	Delete(string) (int, error)
}

type InterHttpClient interface {
	HttpClient
}

type HttpImpl struct {
}

func (this *HttpImpl) Get(url string) ([]byte, error) {

	var (
		resq *http.Response
		err  error
	)

	if resq, err = http.Get(url); err != nil {
		return nil, ErrorHttpNil
	}

	defer resq.Body.Close()

	var (
		resp []byte
	)

	if resp, err = ioutil.ReadAll(resq.Body); err != nil {
		return nil, ErrorByteNil
	}

	return resp, nil

}

func (this *HttpImpl) Post(url string, body map[string]interface{}) ([]byte, error) {

	return nil, nil
}
func (this *HttpImpl) Put(url string, body map[string]interface{}) ([]byte, error) {
	return nil, nil
}
func (this *HttpImpl) Delete(url string) (int, error) {
	return 0, nil
}

func Show(param interface{}) {
	switch param.(type) {
	case *HttpImpl:
		fmt.Println("1")
	default:
		fmt.Println("0")
	}
}

func Say(name interface{}) {
	fmt.Println(name)
}

func main() {
	var httpImpl = &HttpImpl{}
	var httpClient HttpClient
	httpClient = httpImpl
	responseOne, _ := httpClient.Get("https://www.jianshu.com/u/58f0817209aa")
	fmt.Println(string([]byte(responseOne)))

	// interHttpClient

	var interHttpClient InterHttpClient
	interHttpClient = httpImpl
	responseTwo, _ := interHttpClient.Get("http://www.gitbub.com")
	fmt.Println(string([]byte(responseTwo)))

	Show(httpClient)
	Show(interHttpClient)
	// ok..idiom

	if n, ok := httpClient.(*HttpImpl); ok {
		fmt.Println(n)

	}
	if m, ok := interHttpClient.(*HttpImpl); ok {
		fmt.Println(m)
	}

}
