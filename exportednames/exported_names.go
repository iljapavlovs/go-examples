package exportednames

import "fmt"

func someFunctionThatCanNotBeExported() {
	fmt.Println("Not Exported Function")
}

func SomeFunctionThatCanNotBeExported() {
	fmt.Println("This is exported function since it starts with capital letter")
}
