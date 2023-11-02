package main

import (
	"fmt"
	"math/rand"
)

/*
RandBool

	This function returns a random boolean value based on the current time
*/
func RandBool() bool {

	intn := rand.Intn(2)
	fmt.Println(intn)
	return intn == 1
}

// includes min and max in the range
func RandInRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandBool())
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))
	fmt.Println(RandInRange(0, 9))

}
