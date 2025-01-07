package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/usecase/sort"
)

type Copilot interface {
	ListRecipes(sortFunc sort.Yield[recipe.RecipeGroup]) ([]recipe.RecipeGroup, error)
	SaveRecipe(rg recipe.RecipeGroup) error
	SaveRecipeType(rt recipe.RecipeType) error
	SaveGlassType(gt recipe.GlassType) error
	FindRecipeGroup(name string) (recipe.RecipeGroup, error)
	FindRecipeType() (map[string]recipe.RecipeType, error)
	FindGlassType() (map[string]recipe.GlassType, error)
}

func New(deps Dependencies) Copilot {
	return &copilot{
		recipe: recipe.NewSaveLister(deps.Recipe),
	}
}

type Dependencies struct {
	Recipe recipe.Repository
}

type copilot struct {
	recipe recipe.SaveLister
}
