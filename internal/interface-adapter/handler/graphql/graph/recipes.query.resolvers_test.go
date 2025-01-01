package graph

import (
	"context"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_queryResolver_Recipes(t *testing.T) {
	t.Parallel()

	t.Run("will call converter.RecipeGroups with values got from recipe.Service", func(t *testing.T) {
		t.Parallel()

		wantCallConvertRecipeGroupArgs := struct {
			recipeGroups []recipes.RecipeGroup
			recipeTypes  map[string]model.RecipeType
			glassTypes   map[string]model.GlassType
		}{}
		convertOut := []*model.RecipeGroup{
			{Name: "FoundRecipeGroup"},
		}
		want := convertOut

		mockRecipeService := new(MockRecipeService)
		mockRecipeService.On("Find").Return(wantCallConvertRecipeGroupArgs.recipeGroups, nil)
		mockRecipeService.On("FindRecipeType").Return(wantCallConvertRecipeGroupArgs.recipeTypes, nil)
		mockRecipeService.On("FindGlassType").Return(wantCallConvertRecipeGroupArgs.glassTypes, nil)

		mockConverter := new(MockConverter)
		mockConverter.On("RecipeGroups", mock.Anything, mock.Anything, mock.Anything).Return(convertOut)

		r := &queryResolver{&Resolver{Dependencies{
			Orders:         nil,
			Recipes:        mockRecipeService,
			convertToModel: mockConverter,
		}}}

		got, err := r.Recipes(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		mockConverter.AssertCalled(t, "RecipeGroups",
			wantCallConvertRecipeGroupArgs.recipeGroups,
			wantCallConvertRecipeGroupArgs.recipeTypes,
			wantCallConvertRecipeGroupArgs.glassTypes,
		)
	})
}
