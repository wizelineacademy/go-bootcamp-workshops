// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	values := make(chan int)
	shutdown := make(chan struct{})

	// runtime.GOMAXPROCS(0) is used to size the pool based on number of processors
	poolSize := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		// What should you be starting inside this loop?
		// Here's a gift: starts with "g" and ends with "oroutine"
		{

			// Which kind of loop would work for this problem?
			{

				// Generate a random number up to 1000.

				// Use a select to either send the number or receive the shutdown signal.
				{

					// In one case send the random number.

					// In another case receive from the shutdown channel.

				}
			}
		} // Don't forget the values that you need to send to the goroutine's params
	}

	// What are you going to use to store the random numbers?

	// Do you remember how to iterate a channel?
	{

		// What to do if the value is even?

		// And if it's odd?

		// When do you need to break this loop?
	}

	fmt.Println("Receiver sending shutdown signal")
	close(shutdown)

	wg.Wait()

	fmt.Printf("Result count: %d\n", len(nums))
	fmt.Println(nums)
}
