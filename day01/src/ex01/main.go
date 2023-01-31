package main

import (
	"ex01/compare"
	"ex01/entities"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Not enough args")
		return
	}

	var oldFile, newFile string
	for ind, val := range os.Args {
		if val == "--old" && (ind+1) <= (len(os.Args)-1) {
			oldFile = os.Args[ind+1]
		}
		if val == "--new" && (ind+1) <= (len(os.Args)-1) {
			newFile = os.Args[ind+1]
		}
	}
	if oldFile == "" || newFile == "" {
		fmt.Println("I failed to get filenames")
		return
	}

	oldReader, err := initReader(oldFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	newReader, err := initReader(newFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	compare.Check(oldReader, newReader)
}

func initReader(oldFile string) (*entities.DBReader, error) {
	var recipe entities.DBReader
	if strings.HasSuffix(oldFile, ".xml") {
		recipe = &entities.XMLReader{Filename: oldFile}
	} else if strings.HasSuffix(oldFile, ".json") {
		recipe = &entities.JsonReader{Filename: oldFile}
	} else {
		return nil, fmt.Errorf("File extension is wrong")
	}

	recipe.ReadFile()
	return &recipe, nil
}
