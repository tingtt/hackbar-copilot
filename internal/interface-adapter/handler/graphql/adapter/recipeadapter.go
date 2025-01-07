package adapter

type RecipeAdapter interface {
	RecipeAdapterIn
	RecipeAdapterOut
}

func NewRecipeAdapter() RecipeAdapter {
	return &recipeAdapter{NewRecipeAdapterIn(), NewRecipeAdapterOut()}
}

type recipeAdapter struct {
	RecipeAdapterIn
	RecipeAdapterOut
}
