/*
Receive on Quit Channel

How do we know it's finished? Wait for it to tell us it's done: receive on the quit
channel
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	c := simple("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "Bye!"
	fmt.Printf("Joe Says: %q\n", <-quit)
}

func simple(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// Do Nothing.
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()

	return c // Return the channel to the caller.
}

func cleanup() {
	fmt.Println("Cleaned Up")
}
