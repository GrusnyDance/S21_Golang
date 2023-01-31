package main

import (
	"ex00/entities"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
	ReadFile() error
	PrintFile()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough args")
		return
	}

	var recipe DBReader
	filename := os.Args[len(os.Args)-1]
	if strings.HasSuffix(filename, ".xml") {
		recipe = &entities.XMLReader{Filename: filename}
	} else if strings.HasSuffix(filename, ".json") {
		recipe = &entities.JsonReader{Filename: filename}
	} else {
		fmt.Println("File extension is wrong")
		return
	}

	if err := recipe.ReadFile(); err != nil {
		fmt.Println(err)
		return
	}
	recipe.PrintFile()
}
