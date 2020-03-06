package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func reqInit(req *http.Request) {
	req.Header.Add("User-Agent", "tozsy")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", "")
}

func echo() *http.Request {
	req, err := http.NewRequest("GET", "https://www.yuque.com/api/v2/hello", nil)
	if err != nil {
		panic(err)
	}
	reqInit(req)
	return req
}

func userRepos() *http.Request {
	req, err := http.NewRequest("GET", "https://www.yuque.com/api/v2/users/herezsy/repos", nil)
	if err != nil {
		panic(err)
	}
	reqInit(req)
	return req
}

func repos() *http.Request {
	req, err := http.NewRequest("GET", "https://www.yuque.com/api/v2/repos/herezsy/edp0g0", nil)
	if err != nil {
		panic(err)
	}
	reqInit(req)
	return req
}

func docs() *http.Request {
	req, err := http.NewRequest("GET", "https://www.yuque.com/api/v2/repos/herezsy/edp0g0/docs/rw89q0", nil)
	if err != nil {
		panic(err)
	}
	reqInit(req)
	return req
}

func main() {
	client := &http.Client{}

	req := docs()

	rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
