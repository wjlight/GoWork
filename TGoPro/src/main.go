package main

import (
	"fmt"
)

func main() {
	var x int
	x = 10

	printNum(x)
	fmt.Println("HelloWorld")
}

func printNum(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("num = %d", i)
	}
}
