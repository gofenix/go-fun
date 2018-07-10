package main

import (
	"fmt"
	"go-fun/interface/mock"
	"go-fun/interface/queue"
	"go-fun/interface/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is mock"}

	r = real.Retriever{}
	fmt.Println(download(r))

	var q queue.Queue
	q.Push(1)
	q.Push("hello ")
	fmt.Println(q)
}
