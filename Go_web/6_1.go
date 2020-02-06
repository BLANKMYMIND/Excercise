package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// cookie 测试
func testCookie(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 获取名为 test 的 cookie
	c, err := r.Cookie("test")
	// 出错即为空
	if err != nil {
		fmt.Println(err)

		// 注意！ Cookies 的设置必须先于 rw 的输出！！否则设置无效
		// 设置过期时间， MaxAge（以秒作单位） 优先于 Expires
		expiration := time.Now()
		expiration = expiration.Add(time.Second * 60)
		cookie := http.Cookie{Name: "test", Value: "hahaha", Expires: expiration}
		http.SetCookie(rw, &cookie)

		// 按照原来的 cookie 输出到 rw
		for _, cookie := range r.Cookies() {
			fmt.Fprintln(rw, cookie.Name)
		}
		fmt.Fprintln(rw, "There is Not a Cookie named test.")

		// 现有的 cookie
		fmt.Println("Cookie", cookie)
		if c != nil {
			fmt.Println(c.Name, c.Value)
		}
		return
	}
	// 打印 cookie c
	fmt.Println(c.Name, c.Value)

	// 输出到 rw
	fmt.Fprintln(rw, "There is a Cookie named test.")
	fmt.Fprintln(rw, c.Value)
}

func main() {
	// 设置 Handler
	http.HandleFunc("/hello", testCookie)

	// 启动监听
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
