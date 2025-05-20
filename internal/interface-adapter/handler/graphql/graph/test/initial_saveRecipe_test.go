package test

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"hackbar-copilot/internal/utils/sliceutil"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

func InitialSaveRecipes() []IntegrationTestRequest {
	return sliceutil.Map(ParamsSaveRecipe(), func(p graphql.RawParams) IntegrationTestRequest {
		return IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "bartender@example.test",
				Roles:  []string{"bartender"},
				GitHub: &jwtclaims.ClaimsGitHub{ID: "bartender"},
			}),
			body: &p,
		}
	})
}

func ParamsSaveRecipe() []graphql.RawParams {
	params := []graphql.RawParams{}
	for _, variable := range VariablesSaveRecipe {
		params = append(params, graphql.RawParams{
			Query:     QuerySaveRecipe,
			Variables: map[string]any{"input": variable},
		})
	}
	return params
}

var VariablesSaveRecipe = []map[string]any{
	{
		"name":     "Bombay Sapphire",
		"imageURL": "https://example.test/gintonic.webp",
		"replace":  nil,
		"recipes": []map[string]any{
			{
				"name":       "Tonic",
				"category":   "Cocktail",
				"recipeType": map[string]any{"name": "Build"},
				"glassType":  map[string]any{"name": "Collins"},
				"steps": []map[string]any{
					{"material": "Bombay Sapphire", "amount": "45ml"},
					{"description": "Stir"},
					{"material": "Tonic Water", "amount": "full up"},
				},
				"remove": nil,
				"asMenu": nil,
			},
		},
		"remove": nil,
		"asMenu": nil,
	},
	{
		"name":     "Maker's Mark",
		"imageURL": "https://example.test/makersmark.webp",
		"replace":  nil,
		"recipes":  templateInputRecipeGroupRecipesWhisky("Maker's Mark", 800),
		"remove":   nil,
		"asMenu":   map[string]any{},
	},
	{
		"name":     "Phuket Sling",
		"imageURL": "https://example.test/phuketsling.webp",
		"recipes": []map[string]any{
			{
				"name":       "Cocktail",
				"category":   "Cocktail",
				"recipeType": map[string]any{"name": "Build"},
				"glassType":  map[string]any{"name": "Collins"},
				"steps": templateInputStepCocktailLong(
					[]map[string]any{
						{"material": "Peach Liqueur", "amount": "30ml"},
						{"material": "Blue Curacao", "amount": "15ml"},
						{"material": "Grapefruit Juice", "amount": "30ml"},
					},
					"Tonic Water",
				),
				"asMenu": map[string]any{"price": 700},
			},
			{
				"name":       "Mocktail",
				"category":   "NonAlcholic",
				"recipeType": map[string]any{"name": "Build"},
				"glassType":  map[string]any{"name": "Collins"},
				"steps": templateInputStepCocktailLong(
					[]map[string]any{
						{"material": "Peach Syrup", "amount": "20ml"},
						{
							"material": "Blue Curacao Syrup",
							"amount":   "10ml",
						},
						{"material": "Grapefruit Juice", "amount": "30ml"},
					},
					"Tonic Water",
				),
				"asMenu": map[string]any{"price": 700},
			},
		},
		"asMenu": map[string]any{"flavor": "Sweet and Sour"},
	},
}

func templateInputStepCocktailLong(base []map[string]any, mixer string) []map[string]any {
	return append(base,
		map[string]any{"description": "Stir (Well)"},
		map[string]any{"material": mixer, "amount": "Full up"},
	)
}

func templateInputRecipeGroupRecipesWhisky(name string, price int) []map[string]any {
	return []map[string]any{
		{
			"name":       "Rock",
			"category":   "Whisky",
			"recipeType": map[string]any{"name": "Build"},
			"glassType":  map[string]any{"name": "Rock"},
			"steps": []map[string]any{
				{"material": name, "amount": "30ml"},
			},
			"asMenu": map[string]any{"price": price},
		},
		{
			"name":       "Rock (Double)",
			"category":   "Whisky",
			"recipeType": map[string]any{"name": "Build"},
			"glassType":  map[string]any{"name": "Rock"},
			"steps": []map[string]any{
				{"material": name, "amount": "60ml"},
			},
			"asMenu": map[string]any{"price": price * 2},
		},
		{
			"name":       "Soda",
			"category":   "Whisky",
			"recipeType": map[string]any{"name": "Build"},
			"glassType":  map[string]any{"name": "Collins"},
			"steps": templateInputStepCocktailLong(
				[]map[string]any{
					{"material": name, "amount": "30ml"},
				},
				"Soda",
			),
			"asMenu": map[string]any{"price": price},
		},
		{
			"name":       "Straight",
			"category":   "Whisky",
			"recipeType": map[string]any{"name": "Build"},
			"glassType":  map[string]any{"name": "Straight"},
			"steps": []map[string]any{
				{"material": name, "amount": "30ml"},
			},
			"asMenu": map[string]any{"price": price},
		},
		{
			"name":       "Water",
			"category":   "Whisky",
			"recipeType": map[string]any{"name": "Build"},
			"glassType":  map[string]any{"name": "Straight"},
			"steps": templateInputStepCocktailLong(
				[]map[string]any{{
					"material":    name,
					"amount":      "30ml",
					"description": "No ice",
				}},
				"Water",
			),
			"asMenu": nil,
		},
	}
}
