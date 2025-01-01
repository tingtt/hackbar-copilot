package graph

import (
	"context"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_mutationResolver_SaveRecipe(t *testing.T) {
	t.Parallel()

	t.Run("will call recipes.Service.Register with input of arg", func(t *testing.T) {
		t.Parallel()

		input := model.InputRecipeGroup{
			Name: "NewRecipeGroup",
			Recipes: []*model.InputRecipe{{
				Name: "NewRecipe",
				RecipeType: &model.InputRecipeType{
					Name: "NewRecieType",
				},
				GlassType: &model.InputGlassType{
					Name: "NewGlassType",
				},
				Steps: []string{"Step 1", "Step 2"},
			}},
		}

		mockRecipeService := new(MockRecipeService)
		mockRecipeService.On("Register", mock.Anything).Return(recipes.RecipeGroup{}, nil)
		mockRecipeService.On("FindRecipeType").Return(map[string]model.RecipeType{}, nil)
		mockRecipeService.On("FindGlassType").Return(map[string]model.GlassType{}, nil)
		r := &mutationResolver{
			Resolver: &Resolver{
				deps: Dependencies{
					Orders:  nil,
					Recipes: mockRecipeService,
				},
			},
		}

		_, err := r.SaveRecipe(context.Background(), input)

		assert.NoError(t, err)
		mockRecipeService.AssertCalled(t, "Register", input)
	})
}
