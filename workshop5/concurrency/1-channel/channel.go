/*
When the main function executes <-c, it will wait for a value to be sent.

When the boring function executes c <- value, it waits for a receiver to
be ready.

A sender and receiver must both be ready to play their part in the
communication. Otherwise we wait until they are.

Thus channels both communicate and synchronize.

Don't communicate by sharing memory, share memory by communicating.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go simple("simple!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	fmt.Println("You're too simple: I'm leaving.")
}

func simple(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)

		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
