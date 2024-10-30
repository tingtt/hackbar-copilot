package recipes

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type Service interface {
	Register(input model.InputRecipeGroup) (RecipeGroup, error)
	Find() ([]RecipeGroup, error)
	FindRecipeType() (map[string]model.RecipeType, error)
	FindGlassType() (map[string]model.GlassType, error)
}

type Repository interface {
	Find() ([]RecipeGroup, error)
	FindOne(name string) (RecipeGroup, error)
	Save(RecipeGroup) error
	SaveRecipeType(model.RecipeType) error
	SaveGlassType(model.GlassType) error
	FindRecipeType() (map[string]model.RecipeType, error)
	FindGlassType() (map[string]model.GlassType, error)
}

type RecipeGroup struct {
	Name     string
	ImageURL *string
	Recipes  []Recipe
}

type Recipe struct {
	Name  string
	Type  string // build, stir, shake etc.
	Glass string // collins, shot, rock, beer etc.
	Steps []string
}
