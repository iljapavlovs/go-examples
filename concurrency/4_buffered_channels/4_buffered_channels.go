// https://betterprogramming.pub/golang-how-to-implement-concurrency-with-goroutines-channels-2b78b8077984

package main

import "fmt"

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

//Why didn’t the main program print all 3 values? Well, because the channel only has space for one value
//func main() {
//	fmt.Println("We are executing a goroutine")
//	arr := []int{2, 3, 4}
//	ch := make(chan int)
//	go timesThree(arr, ch)
//	time.Sleep(time.Second) // we let the goroutine return all the values
//	result := <-ch
//	fmt.Printf("The result is: %v", result)
//}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4}

	//1. INTIIALIZE BUFFERED CHANNEL - assigning the channel capacity by passing a second parameter to the make function with the number of elements it can get before it’s read
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result: %v \n", <-ch)
	}
}
