package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 即使捕获的原 json 是小写，但转换后仍可自动带入
type Message struct {
	Images   []map[string]interface{}
	Tooltips map[string]interface{}
}

func main() {
	// 在使用命令行代理的情况下会报代理错误：proxyconnect tcp: dial tcp: lookup socks: no such host
	resp, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	r := resp.Body
	if r == nil {
		panic("r is empty")
	}
	j := json.NewDecoder(r)
	fmt.Println(j)
	var m Message
	err = j.Decode(&m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Images[0]["url"])
}
