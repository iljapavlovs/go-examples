package main

import "fmt"

func main() {

	config1 := Config{MaxConcurrentRequests: 0}
	config2 := Config{MaxConcurrentRequests: 0}
	var config3WithPointer *Config
	var config4 Config

	//NOTE - WHEN USING WITH POINTER, THE VALUE IS CHANGED!!!!!!!!
	setConfigWithPointer(&config1)
	fmt.Println("Set Config with Pointer", config1.MaxConcurrentRequests)

	//NOTE - WHEN USING WITHOUT POINTER, THE VALUE IS NOT CHANGED!!!!!!!!
	setConfigWithoutPointer(config2)
	fmt.Println("Set Config WITHOUT Pointer", config2.MaxConcurrentRequests)

	setConfigWithPointer(&config4)
	fmt.Println("Set Config with Pointer", config4.MaxConcurrentRequests)

	//NOTE - FAILS with "invalid memory address or nil pointer dereference" as config3WithPointer pointer is nil
	setConfigWithPointer(config3WithPointer)
	fmt.Println("Set Config with Pointer", config3WithPointer.MaxConcurrentRequests)

}

type Config struct {
	MaxConcurrentRequests int
}

func setConfigWithPointer(config *Config) {
	config.MaxConcurrentRequests = 1
}

func setConfigWithoutPointer(config Config) {
	config.MaxConcurrentRequests = 2
}
