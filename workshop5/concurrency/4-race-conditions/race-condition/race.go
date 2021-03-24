// Golang program to fix the race
// condition using atomic package
package main

import (
	"fmt"
	"sync"
)

// All goroutines will increment variable c
// waitgroup is waiting for the completion
// of program.
var (
	c         int32
	waitgroup sync.WaitGroup
)

func main() {

	// with the help of Add() function add
	// one for each goroutine
	// a count of total 3
	waitgroup.Add(3)

	// increment with the help
	// of increment() function
	go increment("a_long_string-a_long_string-a_long_string-a_long_string-a_long_string-a_long_string")
	go increment("to_cause_a-to_cause_a-to_cause_a-to_cause_a-to_cause_a-to_cause_a-to_cause_a")
	go increment("race_condition-race_condition-race_condition-race_condition-race_condition-race_condition")

	// waiting for completion
	// of goroutines.
	waitgroup.Wait()

	// print the counter
	fmt.Println("Counter:", c)

}

func increment(name string) {
	// Done() function used
	// to tell that it is done.
	defer waitgroup.Done()

	for range name {
		c++
	}
}
