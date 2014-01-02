package main

import (
	"fmt"
)

type Interge int

func (i *Interge) Add(a Interge) {
	*i += a
}

func (i Interge) Less(a Interge) bool {
	return i < a
}

type IInter interface {
	Add(a Interge)
	Less(a Interge) bool
}

var s Interge = 1
var is IInter = &s

func main() {
	fmt.Println(is.Less(1))
}
