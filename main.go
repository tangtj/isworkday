package main

import "net/http"

func main() {
	http.HandleFunc("/today", Req)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("启动异常")
	}
}
