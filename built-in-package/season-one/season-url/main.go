package main

import (
	"fmt"
	"net/url"
)

var urlCollection struct {
	urlOne   string
	urlTwo   string
	urlThree string
	urlFour  string
	urlFive  string
}

func init() {
	urlCollection.urlOne = "https://www.google.com"
	urlCollection.urlTwo = "http://localhost:8887/v1/api/cloud_api/fetcher?email=1156143589@qq.com"
	// https://developer.readsense.cn/docs/retail/retailv2/regions.html#删除区域
	urlCollection.urlThree = "https://developer.readsense.cn/docs/retail/retailv2/regions.html#%E5%88%A0%E9%99%A4%E5%8C%BA%E5%9F%9F"
	urlCollection.urlFour = "https://joe:joepassword@www.email.com/share_info.txt"
	urlCollection.urlFive = "https://godoc.org/net/url#example-Values"
}

func main() {
	OpUrl(urlCollection.urlOne)
	OpUrl(urlCollection.urlTwo)
	OpUrl(urlCollection.urlThree)
	OpUrl(urlCollection.urlFour)
	OpUrl(urlCollection.urlFive)

}

/*
		type URL struct {
		Scheme     string
		Opaque     string    // encoded opaque data
		User       *Userinfo // username and password information
		Host       string    // host or host:port
		Path       string    // path (relative paths may omit leading slash)
		RawPath    string    // encoded path hint (see EscapedPath method)
		ForceQuery bool      // append a query ('?') even if RawQuery is empty
		RawQuery   string    // encoded query values, without '?'
		Fragment   string    // fragment for references, without '#'
	}
*/
func OpUrl(urlString string) {

	URL, _ := url.Parse(urlString)
	fmt.Println("user", URL.User)
	fmt.Println("scheme", URL.Scheme)
	fmt.Println("host", URL.Host)
	fmt.Println("port", URL.Port())
	fmt.Println("rawQuery", URL.RawQuery)
	fmt.Println("rawPath", URL.RawPath)
	fmt.Println("path", URL.Path)
	fmt.Println("forceQuery", URL.ForceQuery)
	fmt.Println("fragment", URL.Fragment)

}

type Url struct {
	Scheme   string
	User     string
	Password string
	Host     string
	Port     string
	Path     string
	Params   map[string][]string
	Fragment string
}
