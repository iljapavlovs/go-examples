package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan struct{})
	go waitForThing(channel)

	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("Waiting")
		case <-channel:
			fmt.Println("Done!")
			return
		case <-time.After(time.Second * 5):
			fmt.Println("Waiting")
			return

			//default:
			//	fmt.Println("DEFAULT!")
		}
	}
}

func waitForThing(channel chan struct{}) {
	time.Sleep(time.Second * 3)
	channel <- struct{}{}
}
