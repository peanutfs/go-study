/**
 * Web服务器
 * 包http通过任何实现了http.Handler的值来响应HTTP请求：
 * package http
 *
 * type Handler interface {
 * 		ServeHTTP(w ResponseWriter, r *Request)
 * }
 * 本例子中，类型Hello实现了http.Handler
 * 访问http://localhost:4000/ 会看到来记程序的问候
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}