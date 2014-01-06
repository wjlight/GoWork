package main

import (
	"fmt"
)

type IWork interface {
	Open() int
}

type Work struct {
	Name string
}

func (w *Work) Open() int {
	return 3
}

func main() {
	var work interface{}
	work = &Work{Name: "wawa"}
	if inis, ok := work.(IWork); ok {
		num := inis.Open()
		fmt.Printf("num", num)
	}

}
