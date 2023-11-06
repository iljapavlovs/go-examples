package main

import (
	"fmt"
	"math/big"
	"slices"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	//This selects a half-open range which includes the first element, but excludes the last one.

	// new slice excluding 0 element
	fmt.Println(s[1:])

	// new slice only including 0 element
	fmt.Println(s[:1])

	slicesContains()

	//s = s[1:4]
	//fmt.Println(s)
	//
	//s = s[:2]
	//fmt.Println(s)
	//
	//s = s[1:]
	//fmt.Println(s)
}

func slicesContains() {
	bigNumA, b := new(big.Int).SetString("92837498273498248274", 10)
	if !b {
		panic("cannot parse big int")
	}
	s := []*big.Int{big.NewInt(-2), big.NewInt(3), big.NewInt(5), big.NewInt(7), bigNumA, big.NewInt(13)}
	fmt.Println(s)

	bigNumB, b := new(big.Int).SetString("92837498273498248274", 10)
	if !b {
		panic("cannot parse big int")
	}
	sliceContains := slices.Contains(s, bigNumB)
	customContains := contains(s, bigNumB)

	fmt.Println("slices.Contains", sliceContains)
	fmt.Println("customContains", customContains)

	isNegative := func(a *big.Int) bool {
		return a.Cmp(big.NewInt(0)) == -1
	}

	containsNegative := slices.ContainsFunc(s, isNegative)
	fmt.Println("containsNegative", containsNegative)

}

func contains(slice []*big.Int, b *big.Int) bool {
	for _, a := range slice {
		if b.Cmp(a) == 0 {
			return true
		}
	}
	return false
}
