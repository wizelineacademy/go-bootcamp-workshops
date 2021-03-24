// Golang program to fix the race
// condition using atomic package
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// All goroutines will increment variable c
// waitgroup is waiting for the completion
// of program.
var (
	count     = SafeCounter{v: make(map[string]int)}
	waitgroup sync.WaitGroup
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

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
	fmt.Println("Counter:", count.v)
}

func increment(name string) {

	// Done() function used
	// to tell that it is done.
	defer waitgroup.Done()

	for range name {

		go count.Inc("somekey")

		// enter thread in the line by line
		runtime.Gosched()
	}
}
