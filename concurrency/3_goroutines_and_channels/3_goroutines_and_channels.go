// https://betterprogramming.pub/golang-how-to-implement-concurrency-with-goroutines-channels-2b78b8077984
package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are executing a goroutine")
	// We have to initialize it with the function make, the keyword chan and the data type between parenthesis

	//GETTING VALUE FROM GO ROUTINE - USE CHANNELS
	//1. INITIALIZE CHANNEL
	ch := make(chan int)

	//2. SEND CHANNEL INTO GO ROUTINE FUNCTION
	go timesThree(3, ch)

	//4. WAIT ANF READ FROM THE CHANNEL AND ASSIGN TO A VARIABLE
	//This doesn’t mean that the main program will wait for the full goroutine to execute, just until the data is served to the channel.
	result := <-ch
	fmt.Printf("The result is: %v", result)
}

// !!!!!!!!!!!!!!!!!!
// Let’s assume we need the result of the operation.
// Then we need to pass the channel as a parameter to the goroutine function so it returns the result with the characters <- for assigning the value.
func timesThree(number int, ch chan int) {
	result := number * 3
	fmt.Println("Inside Go Routine, Result: ", result)

	//3. VALUE FROM GO ROUTINE FUNCTION IS WRITTEN TO THE CHANNEL INSIDE GO ROUTINE
	ch <- result
}
