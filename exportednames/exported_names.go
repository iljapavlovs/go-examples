package exportednames

import (
	"fmt"
	"math"
)

func someFunctionThatCanNotBeExported() {
	fmt.Println("Not Exported Function")
}

func SomeFunctionThatCanBeExported() {
	fmt.Println("This is exported function since it starts with capital letter")
}

func main() {
	// need to be math.Pi, as you cannot export names with lower case
	//fmt.Println(math.pi)
	fmt.Println(math.Pi)

}
