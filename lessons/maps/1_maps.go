package main

import "fmt"

func main() {
	// make function to make a map
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	delete(vertices, "square")

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	fmt.Println(vertices["aaa"]) //will return 0 as such key does not exist

}
