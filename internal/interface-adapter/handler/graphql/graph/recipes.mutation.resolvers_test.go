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

		r := &mutationResolver{&Resolver{Dependencies{
			Orders:         nil,
			Recipes:        mockRecipeService,
			convertToModel: &converter{},
		}}}

		_, err := r.SaveRecipe(context.Background(), input)

		assert.NoError(t, err)
		mockRecipeService.AssertCalled(t, "Register", input)
	})

	t.Run("will call converter.RecipeGroup with values got from recipe.Service", func(t *testing.T) {
		t.Parallel()

		input := model.InputRecipeGroup{Name: "NewRecipeGroup"}
		wantCallConvertRecipeGroupArgs := struct {
			recipeTypes map[string]model.RecipeType
			glassTypes  map[string]model.GlassType
			recipeGroup recipes.RecipeGroup
		}{}
		convertOut := &model.RecipeGroup{Name: input.Name}
		want := convertOut

		mockRecipeService := new(MockRecipeService)
		mockRecipeService.On("Register", mock.Anything).Return(wantCallConvertRecipeGroupArgs.recipeGroup, nil)
		mockRecipeService.On("FindRecipeType").Return(wantCallConvertRecipeGroupArgs.recipeTypes, nil)
		mockRecipeService.On("FindGlassType").Return(wantCallConvertRecipeGroupArgs.glassTypes, nil)

		mockConverter := new(MockConverter)
		mockConverterRecipeGroupFunc := new(MockConverterRecipeGroupFunc)
		mockConverterRecipeGroupFunc.On("Run", mock.Anything).Return(convertOut)
		mockConverter.On("RecipeGroup", mock.Anything, mock.Anything).Return(mockConverterRecipeGroupFunc.Run)

		r := &mutationResolver{&Resolver{Dependencies{
			Orders:         nil,
			Recipes:        mockRecipeService,
			convertToModel: mockConverter,
		}}}

		got, err := r.SaveRecipe(context.Background(), input)

		assert.NoError(t, err)
		assert.Equal(t, want, got)
		mockConverter.AssertCalled(t, "RecipeGroup",
			wantCallConvertRecipeGroupArgs.recipeTypes,
			wantCallConvertRecipeGroupArgs.glassTypes,
		)
		mockConverterRecipeGroupFunc.AssertCalled(t, "Run", wantCallConvertRecipeGroupArgs.recipeGroup)
	})
}
