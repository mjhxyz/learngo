package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 参数: method url body
	request, err := http.NewRequest(http.MethodGet, "http://imooc.com", nil)
	// 设置 header
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	client := http.Client{
		// 拦截重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("重定向:", req)
			return nil // 返回 nil 就是需要重定向
		},
	}
	// 使用默认的Client执行这个请求
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析响应数据
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
