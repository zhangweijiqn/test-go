package main

import (
	"fmt"
	"log"
	"net/http"
)

// test my handler
type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// test add method to handler

// didnt work

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h) // 第二个参数为: handler Handler
	if err != nil {
		log.Fatal(err)
	}
}

/*
包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求
访问 http://localhost:4000/ 会看到来自程序的问候

//类型 Hello 实现了 `http.Handler`, 接口只有一个方法
package http
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/
