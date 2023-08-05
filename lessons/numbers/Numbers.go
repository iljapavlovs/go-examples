package main

import (
	"fmt"
	"math/big"
)

func main() {

	div := big.NewInt(0).Div(big.NewInt(4987528336989286422), big.NewInt(1e18)) //will return 4, there fore  need to use FLoat
	fmt.Println(div)

	quo := new(big.Float).Quo(new(big.Float).SetInt(big.NewInt(4987528336989286422)), big.NewFloat(1e18))
	fmt.Println(quo)
}
