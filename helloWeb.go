package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,world!")
}
func main() {
	http.HandleFunc("/hello", helloHandler) //分发请求——类似URL路由或者URL映射之类的功能
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
