package main

import "fmt"

//The var statement declares a list of variables; as in function argument lists, the type is last.
//
//A var statement can be at package or function level.
var c, python, java bool

//Variables with initializers
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
