/*
Generator: Function that returns a channel

The boring function returns a channel that lets us communicate with the
boring service it provides.

We can have more instances of the service.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := simple("Joe")
	ann := simple("Ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe) // Joe and Ann are blocking each other.
		fmt.Println(<-ann) // waiting for a message to read.
	}

	fmt.Println("You're simple: I'm leaving.")
}

func simple(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
