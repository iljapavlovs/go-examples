package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	//This selects a half-open range which includes the first element, but excludes the last one.

	// new slice excluding 0 element
	fmt.Println(s[1:])

	// new slice only including 0 element
	fmt.Println(s[:1])

	//s = s[1:4]
	//fmt.Println(s)
	//
	//s = s[:2]
	//fmt.Println(s)
	//
	//s = s[1:]
	//fmt.Println(s)
}
