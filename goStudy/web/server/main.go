package main

import (
	"fmt"
	"net/http"
)

func sayHello(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数
	fmt.Fprintf(resp, "%v\n", req.Form)
	fmt.Fprintf(resp, "path:%s\n", req.URL.Path)
	fmt.Fprintf(resp, "scheme:%s\n", req.URL.Scheme)
	fmt.Fprintf(resp, "hello world!\n")
}

func main() {
	http.HandleFunc("/", sayHello) //设置访问路由

	_ = http.ListenAndServe(":9999", nil) //设置监听端口

}
