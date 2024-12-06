package main

import (
	"fmt"
	"os"

	"example.com/greetings/files"
	"github.com/pelletier/go-toml/v2"
)

func main() {

	tomlModel := files.TomlModel{
		Version: 1,
		Name:    "test",
		Tags:    []string{"tag1", "tag2"},
	}

	realFile, err := os.Create("file.toml")
	if err != nil {
		panic(err)
	}

	//creates temp file which is deleted after the program exits
	tempFile, err := os.CreateTemp("", "node_config") //	os.CreateTemp appends random numbers to the file name to ensure uniqueness and security.
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created temp file: %s\n", tempFile.Name())

	nonTempFile, err := os.Create("node_config") //	os.CreateTemp appends random numbers to the file name to ensure uniqueness and security.
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created non-temp file: %s\n", nonTempFile.Name())
	data, err := toml.Marshal(tomlModel)
	if err != nil {
		panic(err)
	}
	_, err = tempFile.WriteString(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote to file: %s\n", tempFile.Name())

	_, err = realFile.WriteString(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Println(tempFile.Name())
	tempfileString, err := os.ReadFile(tempFile.Name())
	if err != nil {
		panic(err)
	}
	fmt.Print(string(tempfileString))

	fmt.Println(realFile.Name())
	realFileString, err := os.ReadFile(realFile.Name())
	if err != nil {
		panic(err)
	}
	fmt.Print(string(realFileString))

}
