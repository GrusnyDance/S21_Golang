package entities

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type JsonReader struct {
	Filename string
	Recipe   *Recipe
}

func (r *JsonReader) ReadFile() error {
	file, err := os.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	recipe := new(Recipe)
	err = json.Unmarshal(file, recipe)
	if err != nil {
		return err
	} else {
		r.Recipe = recipe
	}
	return nil
}

func (r *JsonReader) PrintFile() {
	//recipe := *r.Recipe
	recipeForPrint, err := xml.MarshalIndent(r.Recipe, "", "    ")
	if err != nil {
		fmt.Println("Json is not correct")
		return
	} else {
		fmt.Println(string(recipeForPrint))
	}
}
