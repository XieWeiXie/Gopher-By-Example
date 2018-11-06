package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()
	fmt.Println(r.PostForm)
	for key, value := range r.PostForm {
		fmt.Println(key, value)
	}

}

func sayFuck(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	r.ParseForm()
	fmt.Println(r.MultipartForm)
}
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/hi", sayHi)
	http.HandleFunc("/say", sayFuck)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
1. 动作：GET、POST、PATCH
2. 请求、响应
3. 头部信息

整体包括两部分：
- 请求服务器资源
- 构建 web 服务，自定义处理器
*/
