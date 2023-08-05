// https://betterprogramming.pub/golang-how-to-implement-concurrency-with-goroutines-channels-2b78b8077984

package main

import (
	"fmt"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}

	//if this is removed, then - > fatal error: all goroutines are asleep - deadlock!
	//This happens because the main program tries to receive a value from the channel but there isn’t an active goroutine able to send it.
	close(ch)
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4, 5, 6}
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	for i := range ch {
		fmt.Printf("Result: %v \n", i)
	}
}

//TODO - DIFFERENCE BETWEEN RANGE AND JUST GETTING "<- CHANNEL" IN A LOOP? COMPARE 4 AND 5 !!!!
