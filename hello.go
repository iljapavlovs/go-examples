// first line needs to be a package
// project needs to have MAIN package if you want to execute it
package main

//factored import statement, can be "import "fmt""
import (
	"errors"
	"example.com/greetings/exportednames"

	// functions related to input and output
	"fmt"
	//"go-examples/exportednames"
	"math"
)

// program starts in function MAIN in package MAIN
func main() {
	fmt.Println("Hello World!")

	variables()
	arrays()
	maps()
	loops()
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

	// sqrt() returns 2 values
	// if error occurred, then err will NOT be NIL
	// if error NOT occurred, then err will be NIL
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//structs - collection of fields so you can group things together in a more logical type
	p := person{name: "Jake", age: 23}
	// access Struct fields via dot notation
	fmt.Println(p.age)

	// POINTERS
	i := 7
	// is useless since i is copied by value
	// incrementWithoutPointer gets a copy of i, increments that, but since we are not returning any value, copy is just discarded and original variable is not modified
	incrementWithoutPointer(i)
	fmt.Println(i)

	// generate a pointer - &i, since argument is a pointer
	incrementWithPointer(&i)
	fmt.Println(i)

}

func variables() {
	var x int

	var z int = 7
	// is the same - shorthand
	y := 8

	sum := x + z + y
	fmt.Println(sum)
}

// arrays always have fixed size
func arrays() {
	var a [5]int
	a[2] = 7
	// shorthand for arrays
	b := [5]int{4, 32, 234, 12, 4}

	// in order to have dynamic size of arrays, use slices
	slice := []int{4, 32, 234, 12, 4}
	slice = append(slice, 123)

	fmt.Println(a, b, slice)

}

func maps() {

	// make function to make a map
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	delete(vertices, "square")

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

}

func loops() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iterate through an array/slice

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// iterate through a map

	myMap := make(map[string]string)
	myMap["a"] = "alpha"
	myMap["b"] = "beta"

	for key, value := range myMap {
		fmt.Println("key", key, "value", value)
	}

}

// if error occurred, then err will NOT be NIL
// if error NOT occurred, then err will be NIL
// Go DOES NOT HAVE EXCEPTIONS
func sqrt(x float64) (float64, error) {
	if x < 0 {
		// we could add some additional value if error occurs
		return 0, errors.New("Undefined for negative numbers")
	}
	// if error NOT occurred, then err will be NIL
	return math.Sqrt(x), nil

}

//STRUCT - collection of fields so you can group things together in a more logical type
type person struct {
	name string
	age  int
}

func pointers() {
	i := 7
	// &i - get memmory address,so this gives us a pointer to i
	fmt.Println(&i)

}

func incrementWithoutPointer(x int) {
	x++
}

func incrementWithPointer(x *int) {
	// need to derefference
	// if * not set, then memory address will be incremented instead
	*x++
}
