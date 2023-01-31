package entities

import "encoding/xml"

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		CakeName        string        `xml:"name" json:"name"`
		CakeTime        string        `xml:"stovetime" json:"time"`
		CakeIngredients []Ingredients `xml:"ingredients>item" json:"ingredients"` // child
	} `xml:"cake" json:"cake"`
}

type Ingredients struct {
	IngredientName  string `xml:"itemname" json:"ingredient_name"`
	IngredientCount string `xml:"itemcount" json:"ingredient_count"`
	IngredientUnit  string `xml:"itemunit" json:"ingredient_unit"`
}
