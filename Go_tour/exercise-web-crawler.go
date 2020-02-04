package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	// 若深度超过，返回
	if depth <= 0 {
		return
	}
	// 若记录过，返回
	if cache.Record(url) {
		return
	}
	// 抓取该页内容
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 申明该页探索完毕
	fmt.Printf("found: %s %q\n", url, body)
	// 由该页面下的链接爬向新页面
	if depth-1 > 0 {
		for _, u := range urls { // range 取值的 1 参是 index，2 参是 value 副本
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	// 主函数进程必须阻塞！！！否则协程全死！！！全死！！！
	time.Sleep(time.Second * 3)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

// urls 的访问缓存 结构体
type urlsCache struct {
	urls map[string]bool
	mux  sync.Mutex
}

type lockCache interface {
	Record(url string) bool
}

// urls 的访问缓存 声明
var cache = urlsCache{urls: make(map[string]bool)}

// urls 缓存函数
func (c *urlsCache) Record(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.urls[url]; ok {
		return true
	}
	c.urls[url] = true
	return false
}

type fakeResult struct {
	body string
	urls []string
}

// fetcher 的 Fetch 方法
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// 简明判断值存在并进一步
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	// 页面不存在时，抛出错误
	return "", nil, fmt.Errorf("not found: %s", url)
}

// 声明 fetcher 变量
// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
