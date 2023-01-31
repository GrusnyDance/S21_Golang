package entities

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type XMLReader struct {
	Filename string
	Recipe   *Recipe
}

func (r *XMLReader) ReadFile() error {
	file, err := os.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	recipe := new(Recipe)
	err = xml.Unmarshal(file, recipe)
	if err != nil {
		return err
	} else {
		r.Recipe = recipe
	}
	return nil
}

func (r *XMLReader) PrintFile() {
	//recipe := *r.Recipe
	recipeForPrint, err := json.MarshalIndent(r.Recipe, "", "    ")
	if err != nil {
		fmt.Println("XML is not correct")
		return
	} else {
		fmt.Println(string(recipeForPrint))
	}
}
