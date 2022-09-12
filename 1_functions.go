package main

import (
	//PACKAGES! - last name - is the name of the package
	"example.com/greetings/exportednames"
	"fmt"
)

func main() {
	//FUNCTIONS
	// exported names from a diff package
	exportednames.SomeFunctionThatCanBeExported()
	//Notice that the type comes after the variable name
	fmt.Println(sum(3, 5))
	//When two or more consecutive named function parameters share a type,
	//you can omit the type from all but the last.
	fmt.Println(add(3, 5))
	// EXAMPLE OF A FUNCTION RETURNING MULTIPLE RESULTS
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//named return values in a function
	fmt.Println(split(17))

}

// FUNCTIONS - Notice that the type comes after the variable name.
func sum(x int, y int) int {
	return x + y
}

//When two or more consecutive named function parameters share a type,
//you can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

// functions can return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

//Named return values
//Go's return values may be named. If so, they are treated as variables defined at the top of the function.
//
//These names should be used to document the meaning of the return values.!!!
//
//A return statement without arguments returns the named return values. This is known as a "naked" return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
