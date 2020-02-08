package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 作静态服务器时，pattern 必须以 / 结尾，同时，prefix 必须为 pattern 子集
	// StripPrefix 是必要的，他将过滤 url 中的 static 字段
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./test"))))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
