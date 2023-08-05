package main

import (
	"fmt"
	"time"
)

//https://go.dev/tour/concurrency/1

func main() {

	// A goroutine is a lightweight thread managed by the Go runtime.
	//starts a new goroutine running
	go say("GO ROUTINE")

	say("NO GO ROUTINE")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.
//
//Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)
