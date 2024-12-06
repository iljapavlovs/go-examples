package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	div := big.NewInt(0).Div(big.NewInt(4987528336989286422), big.NewInt(1e18)) //will return 4, there fore  need to use FLoat
	fmt.Println(div)

	quo := new(big.Float).Quo(new(big.Float).SetInt(big.NewInt(4987528336989286422)), big.NewFloat(1e18))
	fmt.Println(quo)

	//fmt.Println(big.NewInt(1e19)) - cannot use 1e19 (untyped float constant 1e+19) as int64 value in argument to big.NewInt (truncated)

	fmt.Println(new(big.Int).SetString("1e19", 10))

	//fmt.Println(big.NewInt(10000000000000000000)) - cannot use 10000000000000000000 (untyped int constant) as int64 value in argument to big.NewInt (overflows)

	setString, _ := new(big.Int).SetString("10_000_000_000_000_000_000", 10)

	fmt.Println(setString.String())

	x := big.NewInt(1e18)
	y := big.NewInt(10)
	z := x.Mul(x, y)

	fmt.Println(z)

	sub := new(big.Int).Sub(big.NewInt(12), big.NewInt(11))
	fmt.Println(sub)

	fmt.Println("6e16", int64(6e16))

	convertBigFloatToBigInt()

	convertBigIntToBigFloat()

	gas := 2_900_000
	result := gas / 500_000
	fmt.Println("result", result)

}

func convertBigFloatToBigInt() {
	result := new(big.Int)

	//NEED TO SET .SetPrec(100) to get the correct result
	//https://stackoverflow.com/questions/52160754/strange-loss-of-precision-multiplying-big-float
	//https://stackoverflow.com/questions/56133476/precision-issue-of-golang-big-float
	lowResult, accuracyForLow := new(big.Float).Mul(big.NewFloat(0.000000000000000001), big.NewFloat(1e18)).Int(new(big.Int))
	highResult, accuracyForHigh := new(big.Float).SetPrec(100).Mul(big.NewFloat(10000000000000000), big.NewFloat(1e18)).Int(new(big.Int))

	fmt.Println("result", result)
	fmt.Println("accuracyForLow", accuracyForLow)

	fmt.Println("lowResult", lowResult)
	fmt.Println("highResult", highResult)
	fmt.Println("highResult", highResult.String())

	fmt.Println("accuracyForHigh", accuracyForHigh)
	//99999999999999991611392
}

func convertBigIntToBigFloat() {
	bigFloatNumber := new(big.Float).SetInt(big.NewInt(1e18))
	fmt.Println("convertBigIntToBigFloat", bigFloatNumber.String())

	weiToEther := WeiToEther(big.NewInt(1e18))
	fmt.Println("weiToEther", weiToEther.String())

	etherToWei := EtherToWei(big.NewFloat(1.5))
	fmt.Println("etherToWei", etherToWei.String())
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(1e18))
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(1e18))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}
