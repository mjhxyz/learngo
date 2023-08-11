package main

import (
	"fmt"
	"learngo/advance/retriever/mock"
	"learngo/advance/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://z.cn"

func download(r Retriever) string {
	return r.Get(url)
}

func post(p Poster) {
	p.Post(url,
		map[string]string{
			"name": "mao",
			"lang": "go",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked z.cn",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{Contents: "this is a fake!!!"}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mao",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion 获取具体的类型
	// 类似于强制类型转换，如果转不过来，则会报错
	// 可以加一个 ok 参数接收
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&mockRetriever))
	fmt.Println(&mockRetriever)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
