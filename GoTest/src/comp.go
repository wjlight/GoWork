package main

import (
	"fmt"
)

type Entry struct {
	Name string
	Addr string
}

func NewEntry() *Entry {
	return &Entry{Name: "wawa", Addr: "nit"}
}

type IPoint interface {
	Area() int
}

func (e *Entry) Area() int {
	return 2
}

func CheckIsImplInterface(something interface{}) bool {
	t, ok := something.(IPoint)
	fmt.Println("t:", t, " ok:", ok)
	return ok
}

func main() {
	CheckIsImplInterface(&Entry{})
}

func CompField() {
	ar := [...]string{"no error"}
	ma := map[int]string{1: "no error"}
	fmt.Println("ar:", ar, "  ma:", ma[1])
	en := NewEntry()
	fmt.Println("en:", en.Addr)
}
