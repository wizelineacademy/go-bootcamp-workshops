/*
Daisy-Chain

1 ->    2          3          4          5          6          7
     Gopher1 -> Gopher2 -> Gopher3 -> Gopher4 -> Gopher5 -> Gopher6
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	const n = 100000
	rightmost := make(chan int)

	var left chan int
	right := rightmost

	for i := 0; i < n; i++ {
		left = make(chan int)
		go pass(left, right)
		right = left
	}

	go func(c chan int) {
		c <- 1
	}(left)

	fmt.Printf("Final Value: %d\n", <-rightmost)
	fmt.Printf("This took: %f seconds\n", time.Since(start).Seconds())
}

func pass(left, right chan int) {
	value := <-left
	right <- 1 + value
}
