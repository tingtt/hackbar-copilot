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

func Test_queryResolver_Recipes(t *testing.T) {
	t.Parallel()
}

func Test_convertRecipeGroup(t *testing.T) {
	t.Parallel()

	type args struct {
		recipeTypes map[string]model.RecipeType
		glassTypes  map[string]model.GlassType
		recipeGroup recipes.RecipeGroup
	}
	tests := []struct {
		name string
		args args
		want model.RecipeGroup
	}{
		{
			name: "will return merged RecipeGroup contains name and image url",
			args: args{
				recipeGroup: recipes.RecipeGroup{
					Name: "Phuket Sling",
					ImageURL: func() *string {
						text := "https://example.com/path/to/image"
						return &text
					}(),
					Recipes: []recipes.Recipe{},
				},
			},
			want: model.RecipeGroup{
				Name: "Phuket Sling",
				ImageURL: func() *string {
					text := "https://example.com/path/to/image"
					return &text
				}(),
				Recipes: []*model.Recipe{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := convertRecipeGroup(tt.args.recipeTypes, tt.args.glassTypes)(tt.args.recipeGroup)

			assert.Equal(t, *got, tt.want)
		})
	}
}

func Test_convertRecipe(t *testing.T) {
	t.Parallel()

	type args struct {
		recipeTypes map[string]model.RecipeType
		glassTypes  map[string]model.GlassType
		recipe      recipes.Recipe
	}
	tests := []struct {
		name string
		args args
		want model.Recipe
	}{
		{
			name: "will return merged Recipe and match recipeType and GlassType",
			args: args{
				recipeTypes: map[string]model.RecipeType{
					"build": {
						Name: "build",
						Description: func() *string {
							text := "build description"
							return &text
						}(),
					},
				},
				glassTypes: map[string]model.GlassType{
					"collins": {
						Name: "collins",
						ImageURL: func() *string {
							text := "https://example.com/path/to/image"
							return &text
						}(),
						Description: func() *string {
							text := "collins glass description"
							return &text
						}(),
					},
				},
				recipe: recipes.Recipe{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
				},
			},
			want: model.Recipe{
				Name: "Cocktail",
				Type: &model.RecipeType{
					Name: "build",
					Description: func() *string {
						text := "build description"
						return &text
					}(),
				},
				Glass: &model.GlassType{
					Name: "collins",
					ImageURL: func() *string {
						text := "https://example.com/path/to/image"
						return &text
					}(),
					Description: func() *string {
						text := "collins glass description"
						return &text
					}(),
				},
				Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
			},
		},
		{
			name: "will return merged Recipe and don't match recipeType and GlassType",
			args: args{
				recipeTypes: map[string]model.RecipeType{},
				glassTypes:  map[string]model.GlassType{},
				recipe: recipes.Recipe{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
				},
			},
			want: model.Recipe{
				Name:  "Cocktail",
				Type:  nil,
				Glass: nil,
				Steps: []string{"Peach liqueur 30ml", "Blue curacao 15ml", "Grapefruit juice 30ml", "Stir", "Tonic water - Full up"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := convertRecipe(tt.args.recipeTypes, tt.args.glassTypes)(tt.args.recipe)

			assert.Equal(t, *got, tt.want)
		})
	}
}

func TestResolver_Mutation(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, (&Resolver{}).Mutation())
	})
}

func TestResolver_Query(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, (&Resolver{}).Query())
	})
}
