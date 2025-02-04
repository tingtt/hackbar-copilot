package menuadapter

import (
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	recipeadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func Test_menuAdapterOut_MenuGroups(t *testing.T) {
	t.Parallel()

	t.Run("will return adapted menu groups", func(t *testing.T) {
		t.Parallel()

		argRecipeGroups := recipeadapter.NewOutputAdapter().RecipeGroups(
			recipetest.ExampleRecipeGroups,
			recipetest.ExampleRecipeTypesMap,
			recipetest.ExampleGlassTypesMap,
		)
		argMenuGroups := menutest.ExampleGroups
		var want []*model.MenuGroup = []*model.MenuGroup{
			{
				Name:     "Phuket Sling",
				ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
				Flavor:   ptr("Sweet"),
				Items: []*model.MenuItem{
					{
						Name:       "Cocktail",
						ImageURL:   ptr("https://example.com/path/to/image/phuket-sling/cocktail"),
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name: "Cocktail",
							Type: &model.RecipeType{
								Name:        "build",
								Description: ptr("build description"),
							},
							Glass: &model.GlassType{
								Name:        "collins",
								ImageURL:    ptr("https://example.com/path/to/image/collins"),
								Description: ptr("collins glass description"),
							},
							Steps: []*model.Step{
								{
									Material: ptr("Peach liqueur"),
									Amount:   ptr("30ml"),
								},
								{
									Material: ptr("Blue curacao"),
									Amount:   ptr("15ml"),
								},
								{
									Material: ptr("Grapefruit juice"),
									Amount:   ptr("30ml"),
								},
								{
									Description: ptr("Stir"),
								},
								{
									Material: ptr("Tonic water"),
									Amount:   ptr("Full up"),
								},
							},
						},
					},
					{
						Name:       "Mocktail",
						ImageURL:   ptr("https://example.com/path/to/image/phuket-sling/mocktail"),
						Materials:  []string{"Peach syrup", "Blue curacao syrup", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						PriceYen:   500,
						Recipe:     nil,
					},
				},
				MinPriceYen: 500,
			},
			{
				Name:     "Passoamoni",
				ImageURL: ptr("https://example.com/path/to/image/passoamoni"),
				Flavor:   ptr("Fruity"),
				Items: []*model.MenuItem{
					{
						Name:       "Cocktail",
						ImageURL:   ptr("https://example.com/path/to/image/passoamoni"),
						Materials:  []string{"Passoa", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name: "Cocktail",
							Type: &model.RecipeType{
								Name:        "build",
								Description: ptr("build description"),
							},
							Glass: &model.GlassType{
								Name:        "collins",
								ImageURL:    ptr("https://example.com/path/to/image/collins"),
								Description: ptr("collins glass description"),
							},
							Steps: []*model.Step{
								{
									Material: ptr("Passoa"),
									Amount:   ptr("45ml"),
								},
								{
									Material: ptr("Grapefruit juice"),
									Amount:   ptr("30ml"),
								},
								{
									Description: ptr("Stir"),
								},
								{
									Material: ptr("Tonic water"),
									Amount:   ptr("Full up"),
								},
							},
						},
					},
				},
				MinPriceYen: 700,
			},
			{
				Name:     "Blue Devil",
				ImageURL: ptr("https://example.com/path/to/image/blue-devil"),
				Flavor:   ptr("Medium sweet and dry"),
				Items: []*model.MenuItem{
					{
						Name:       "Cocktail",
						ImageURL:   ptr("https://example.com/path/to/image/blue-devil"),
						Materials:  []string{"Gin", "Blue curacao", "Lemon juice"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name: "Cocktail",
							Type: &model.RecipeType{
								Name:        "shake",
								Description: ptr("shake description"),
							},
							Glass: &model.GlassType{
								Name:        "cocktail",
								ImageURL:    ptr("https://example.com/path/to/image/cocktail"),
								Description: ptr("cocktail glass description"),
							},
							Steps: []*model.Step{
								{
									Description: ptr("Chill shaker and glass."),
								},
								{
									Description: ptr("Put ingredients in a shaker."),
								},
								{
									Material: ptr("Gin"),
									Amount:   ptr("30ml"),
								},
								{
									Material: ptr("Blue curacao"),
									Amount:   ptr("15ml"),
								},
								{
									Material: ptr("Lemon juice"),
									Amount:   ptr("15ml"),
								},
								{
									Description: ptr("Put ice in a shaker."),
								},
								{
									Description: ptr("Shake."),
								},
								{
									Description: ptr("Pour into a glass."),
								},
							},
						},
					},
				},
				MinPriceYen: 700,
			},
		}

		adapter := &outputAdapter{}
		got := adapter.MenuGroups(argMenuGroups, argRecipeGroups)
		assert.Equal(t, want, got)
	})
}
