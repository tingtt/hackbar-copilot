package graph

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/recipes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_converter_RecipeGroups(t *testing.T) {
	t.Parallel()

	type args struct {
		recipeTypes map[string]model.RecipeType
		glassTypes  map[string]model.GlassType
		recipeGroup []recipes.RecipeGroup
	}
	tests := []struct {
		name string
		args args
		want []*model.RecipeGroup
	}{
		{
			name: "will return merged RecipeGroup contains name and image url",
			args: args{
				recipeGroup: []recipes.RecipeGroup{{
					Name: "Phuket Sling",
					ImageURL: func() *string {
						text := "https://example.com/path/to/image"
						return &text
					}(),
					Recipes: []recipes.Recipe{},
				}},
			},
			want: []*model.RecipeGroup{{
				Name: "Phuket Sling",
				ImageURL: func() *string {
					text := "https://example.com/path/to/image"
					return &text
				}(),
				Recipes: []*model.Recipe{},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			convertToModel := &converter{}
			got := convertToModel.RecipeGroups(tt.args.recipeGroup, tt.args.recipeTypes, tt.args.glassTypes)

			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_converter_recipe(t *testing.T) {
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

			convertToModel := &converter{}
			got := convertToModel.recipe(tt.args.recipeTypes, tt.args.glassTypes)(tt.args.recipe)

			assert.Equal(t, *got, tt.want)
		})
	}
}
