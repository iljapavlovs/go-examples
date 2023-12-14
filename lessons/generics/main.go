package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Add(x, y int) int {
	return x + y
}

func AddFloat(x, y float64) float64 {
	return x + y
}

// problem - there are a lot of types of int, so this is not elegant
func AddWithGeneric[T int | int8 | int16 | float64 | float32](x, y T) T {
	return x + y
}

///////////////////////////////////////////

type Num interface {
	int | int8 | int16 | float64 | float32
}

// still too verbose
func AddWithGenericAndInterface[T Num](x, y T) T {
	return x + y
}

// /////////////////////////////////////////
// note - need to manually import constraints package
func AddWithContraints[T constraints.Ordered](x, y T) T {
	return x + y
}

// /////////////////////////////////////////
type UserID int

// tilda - allows you to use an alias for this type
func AddWithTilda[T ~int | int8 | int16 | float64 | float32](x, y T) T {
	return x + y
}

// /////////////////////////////////////////

type CustomData interface {
	constraints.Ordered | []byte | string //can be anyhting
}
type User[T CustomData] struct {
	ID   int
	Name string
	//Data interface{} // need to avoid using interface since you will need to cast it and compiler will not help you
	Data T
}

func main() {
	resultInt := Add(1, 2)
	resultFloat := AddFloat(1.1, 2.2)

	addWithGeneric := AddWithGeneric(1, 2)

	addWithGenericAndInterface := AddWithGenericAndInterface(1, 2)

	addWithContraints := AddWithContraints(1, 2)

	a := UserID(1)
	b := UserID(2)
	addWithTilda := AddWithTilda(a, b)

	userWithByte := User[[]byte]{ID: 1, Name: "John", Data: []byte("some data")}
	userWithString := User[string]{ID: 1, Name: "John", Data: "some data"}
	userWithInt := User[int]{ID: 1, Name: "John", Data: 1231231231}

	fmt.Println(resultInt)
	fmt.Println(resultFloat)
	fmt.Println(addWithGeneric)
	fmt.Println(addWithGenericAndInterface)
	fmt.Println(addWithContraints)
	fmt.Println(addWithTilda)
	fmt.Println(userWithByte)
	fmt.Println(userWithString)
	fmt.Println(userWithInt)

}
