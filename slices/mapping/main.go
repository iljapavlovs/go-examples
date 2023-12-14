package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	values := MapValues([]int{1, 2, 3}, func(x int) int { return x + 1 })
	fmt.Println(values)

}

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, value := range values {
		newValues = append(newValues, mapFunc(value))
	}
	return newValues
}
