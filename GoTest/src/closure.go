package main

import (
	"fmt"
)

func main() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d \n", i, j)
		}
	}() //花括号后直接跟参数列表表示函数调用

	a() //i, j: 10, 5   匿名函数形成了闭包

	j *= 2

	a() //i, j: 10, 10
}
