package compare

import (
	"ex01/entities"
	"fmt"
)

func Check(old *entities.DBReader, new *entities.DBReader) {
	oldRecipe := (*old).GetRecipes()
	newRecipe := (*new).GetRecipes()
	cakesCompare(oldRecipe, newRecipe)
}

func ingredientDetailsCompare(oldIngredient, newIngredient *entities.Ingredients, cakeName string) {
	if oldIngredient.IngredientCount != newIngredient.IngredientCount {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			oldIngredient.IngredientName,
			cakeName,
			newIngredient.IngredientCount,
			oldIngredient.IngredientCount)
	}
	if oldIngredient.IngredientUnit != newIngredient.IngredientUnit {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			oldIngredient.IngredientName,
			cakeName,
			newIngredient.IngredientUnit,
			oldIngredient.IngredientUnit)
	}
	if oldIngredient.IngredientUnit == "" && newIngredient.IngredientUnit != "" {
		fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			newIngredient.IngredientUnit,
			newIngredient.IngredientName,
			cakeName)
	}
	if newIngredient.IngredientUnit == "" && oldIngredient.IngredientUnit != "" {
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
			oldIngredient.IngredientUnit,
			oldIngredient.IngredientName,
			cakeName)
	}
}

func ingredientCompare(old, new *[]entities.Ingredients, cakeName string) {
	for _, valueOld := range *old {
		delete := true
		for _, valueNew := range *new {
			if valueOld.IngredientName == valueNew.IngredientName {
				ingredientDetailsCompare(&valueOld, &valueNew, cakeName)
				delete = false
				break
			}
		}
		if delete {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", valueOld.IngredientName, cakeName)
		}
	}

	for _, valueNew := range *new {
		add := true
		for _, valueOld := range *old {
			if valueNew.IngredientName == valueOld.IngredientName {
				add = false
				break
			}
		}
		if add {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", valueNew.IngredientName, cakeName)
		}
	}
}

func cakesCompare(oldRecipe, newRecipe *entities.Recipe) {
	for _, valOld := range (*oldRecipe).Cake {
		delete := true
		for _, valNew := range (*newRecipe).Cake {
			if valOld.CakeName == valNew.CakeName {
				if valOld.CakeTime != valNew.CakeTime {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
						valOld.CakeName,
						valNew.CakeTime,
						valOld.CakeTime)
				}
				delete = false
				break
			}
		}
		if delete {
			fmt.Printf("REMOVED cake \"%s\"\n", valOld.CakeName)
		}
	}

	for _, valNew := range (*newRecipe).Cake {
		add := true
		for _, valOld := range (*oldRecipe).Cake {
			if valNew.CakeName == valOld.CakeName {
				ingredientCompare(&valOld.CakeIngredients, &valNew.CakeIngredients, valNew.CakeName)
				add = false
				break
			}
		}
		if add {
			fmt.Printf("ADDED cake \"%s\"\n", valNew.CakeName)
		}
	}
}
