package main

import (
	"fmt"
)

// Sendthe sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
		// fmt.Println("hel", i)
	}
}

// Copythe values from channel 'in' to channel 'out',
//removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		fmt.Println("Filter", i, "  prime:", prime)
		if i%prime != 0 {
			fmt.Println("Filter Write")
			out <- i
		}
	}
	fmt.Println("Filter Over")
}

func main() {
	ch := make(chan int)

	go Generate(ch)

	// prime := <-ch
	// print(prime, "\n")

	// prime = <-ch
	// print(prime, "\n")
	for i := 0; i < 3; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		fmt.Println("go method go on...")
		// print(<-ch1, "\n")
		ch = ch1
	}
}
