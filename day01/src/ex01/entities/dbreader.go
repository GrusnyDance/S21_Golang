package entities

type DBReader interface {
	ReadFile() error
	PrintFile()
	GetRecipes() *Recipe
}
