package menuadapter

import (
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	recipeadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/recipe"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_menuAdapterOut_MenuItems(t *testing.T) {
	t.Parallel()

	t.Run("will return adapted menu items", func(t *testing.T) {
		t.Parallel()

		argRecipeGroups := recipeadapter.NewOutputAdapter().RecipeGroups(
			recipetest.ExampleRecipeGroups,
			recipetest.ExampleRecipeTypesMap,
			recipetest.ExampleGlassTypesMap,
		)
		argMenuItems := menutest.ExampleItems
		var want []*model.MenuItem = []*model.MenuItem{
			{
				Name:     "Phuket Sling",
				ImageURL: new("https://example.com/path/to/image/phuket-sling"),
				Flavor:   new("Sweet"),
				Options: []*model.MenuItemOption{
					{
						Name:       "Cocktail",
						Category:   "Cocktail",
						ImageURL:   new("https://example.com/path/to/image/phuket-sling/cocktail"),
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name:     "Cocktail",
							Category: "Cocktail",
							Type: &model.RecipeType{
								Name:        "build",
								Description: new("build description"),
							},
							Glass: &model.GlassType{
								Name:        "collins",
								ImageURL:    new("https://example.com/path/to/image/collins"),
								Description: new("collins glass description"),
							},
							Steps: []*model.Step{
								{
									Material: new("Peach liqueur"),
									Amount:   new("30ml"),
								},
								{
									Material: new("Blue curacao"),
									Amount:   new("15ml"),
								},
								{
									Material: new("Grapefruit juice"),
									Amount:   new("30ml"),
								},
								{
									Description: new("Stir"),
								},
								{
									Material: new("Tonic water"),
									Amount:   new("Full up"),
								},
							},
						},
					},
					{
						Name:       "Mocktail",
						Category:   "Mocktail",
						ImageURL:   new("https://example.com/path/to/image/phuket-sling/mocktail"),
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
				ImageURL: new("https://example.com/path/to/image/passoamoni"),
				Flavor:   new("Fruity"),
				Options: []*model.MenuItemOption{
					{
						Name:       "Cocktail",
						Category:   "Cocktail",
						ImageURL:   new("https://example.com/path/to/image/passoamoni"),
						Materials:  []string{"Passoa", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name:     "Cocktail",
							Category: "Cocktail",
							Type: &model.RecipeType{
								Name:        "build",
								Description: new("build description"),
							},
							Glass: &model.GlassType{
								Name:        "collins",
								ImageURL:    new("https://example.com/path/to/image/collins"),
								Description: new("collins glass description"),
							},
							Steps: []*model.Step{
								{
									Material: new("Passoa"),
									Amount:   new("45ml"),
								},
								{
									Material: new("Grapefruit juice"),
									Amount:   new("30ml"),
								},
								{
									Description: new("Stir"),
								},
								{
									Material: new("Tonic water"),
									Amount:   new("Full up"),
								},
							},
						},
					},
				},
				MinPriceYen: 700,
			},
			{
				Name:     "Blue Devil",
				ImageURL: new("https://example.com/path/to/image/blue-devil"),
				Flavor:   new("Medium sweet and dry"),
				Options: []*model.MenuItemOption{
					{
						Name:       "Cocktail",
						Category:   "Cocktail",
						ImageURL:   new("https://example.com/path/to/image/blue-devil"),
						Materials:  []string{"Gin", "Blue curacao", "Lemon juice"},
						OutOfStock: false,
						PriceYen:   700,
						Recipe: &model.Recipe{
							Name:     "Cocktail",
							Category: "Cocktail",
							Type: &model.RecipeType{
								Name:        "shake",
								Description: new("shake description"),
							},
							Glass: &model.GlassType{
								Name:        "cocktail",
								ImageURL:    new("https://example.com/path/to/image/cocktail"),
								Description: new("cocktail glass description"),
							},
							Steps: []*model.Step{
								{
									Description: new("Chill shaker and glass."),
								},
								{
									Description: new("Put ingredients in a shaker."),
								},
								{
									Material: new("Gin"),
									Amount:   new("30ml"),
								},
								{
									Material: new("Blue curacao"),
									Amount:   new("15ml"),
								},
								{
									Material: new("Lemon juice"),
									Amount:   new("15ml"),
								},
								{
									Description: new("Put ice in a shaker."),
								},
								{
									Description: new("Shake."),
								},
								{
									Description: new("Pour into a glass."),
								},
							},
						},
					},
				},
				MinPriceYen: 700,
			},
		}

		adapter := &outputAdapter{}
		got := adapter.MenuItems(argMenuItems, argRecipeGroups)
		assert.Equal(t, want, got)
	})
}
