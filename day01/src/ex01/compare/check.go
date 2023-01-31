package compare

import (
	"ex01/entities"
	"fmt"
	"github.com/r3labs/diff/v3"
	"strconv"
)

var KeyValForPrint map[string]string

func Check(old *entities.DBReader, new *entities.DBReader) {
	KeyValForPrint = map[string]string{
		"Cake":            "cake",
		"CakeTime":        "cooking time",
		"CakeIngredients": "ingredient",
		"IngredientCount": "unit count",
		"IngredientUnit":  "unit",
	}
	changelog, _ := diff.Diff((*old).GetRecipes(), (*new).GetRecipes())

	for _, val := range changelog {
		if val.Path[len(val.Path)-1] == "Filename" || val.Path[len(val.Path)-1] == "Local" {
			continue
		}
		fmt.Println(val)
		if val.Type == "create" {
			addedItem(&val, old)
		} else if val.Type == "delete" {
			removedItem(&val, old)
		} else if val.Type == "update" {
			changedItem(&val, old)
		}
	}
}

func addedItem(item *diff.Change, old *entities.DBReader) {
	fmt.Printf("ADDED %s\n", item.Path[len(item.Path)-1])

}

func removedItem(item *diff.Change, old *entities.DBReader) {
	keyHelper := item.Path[len(item.Path)-1]
	if keyHelper == "IngredientName" {
		fmt.Printf("REMOVED ")
	} else {
		fmt.Printf("REMOVED %s ", KeyValForPrint[item.Path[len(item.Path)-1]])
	}

	recipe := (*old).GetRecipes()
	cake := getCake(item.Path, recipe)
	for i := len(item.Path) - 2; i >= 1; i = i - 2 {
		if i == 1 {
			fmt.Printf("for %s \"%s\" ", "cake", cake)
		} else if i == 3 {
			ingredient := getIngredient(item.Path, recipe)
			if keyHelper != "IngredientName" {
				fmt.Printf("for %s \"%s\" ", KeyValForPrint[item.Path[2]], ingredient)
			} else {
				fmt.Printf("%s \"%s\" ", KeyValForPrint[item.Path[2]], ingredient)
			}
		}
	}
	fmt.Printf("\n")
}

func changedItem(item *diff.Change, old *entities.DBReader) {
	fmt.Printf("CHANGED %s ", KeyValForPrint[item.Path[len(item.Path)-1]])
	recipe := (*old).GetRecipes()
	cake := getCake(item.Path, recipe)
	for i := len(item.Path) - 2; i >= 1; i = i - 2 {
		if i == 1 {
			fmt.Printf("for %s \"%s\" ", "cake", cake)
		} else if i == 3 {
			ingredient := getIngredient(item.Path, recipe)
			fmt.Printf("for %s \"%s\" ", KeyValForPrint[item.Path[2]], ingredient)
		}
	}

	fmt.Printf("- \"%s\" instead of \"%s\"\n", item.To, item.From)
}

func getCake(path []string, recipe *entities.Recipe) string {
	if len(path) >= 2 {
		ind, _ := strconv.Atoi(path[1])
		return recipe.Cake[ind].CakeName
	}
	return ""
}

func getIngredient(path []string, recipe *entities.Recipe) string {
	cakeInd, _ := strconv.Atoi(path[1])
	ingredientInd, _ := strconv.Atoi(path[3])
	return recipe.Cake[cakeInd].CakeIngredients[ingredientInd].IngredientName
}
