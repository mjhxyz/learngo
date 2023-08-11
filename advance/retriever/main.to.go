package main

import (
	"fmt"
	"learngo/advance/retriever/mock"
	"learngo/advance/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://z.cn")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake!!!"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
