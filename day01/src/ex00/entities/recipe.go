package entities

import "encoding/xml"

type Recipe struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			Name  string `xml:"itemname" json:"ingredient_name"`
			Count string `xml:"itemcount" json:"ingredient_count"`
			Unit  string `xml:"itemunit,omitempty" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"` // child
	} `xml:"cake" json:"cake"`
}
