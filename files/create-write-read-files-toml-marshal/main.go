package main

import (
	"example.com/greetings/files"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
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
	tempFile, err := os.CreateTemp("", "node_config")
	if err != nil {
		panic(err)
	}
	data, err := toml.Marshal(tomlModel)
	if err != nil {
		panic(err)
	}
	_, err = tempFile.WriteString(string(data))
	if err != nil {
		panic(err)
	}

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
