package test

import (
	"context"
	graphqlhandler "hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var getRecipeTests = []IntegrationTest{
	{
		name:   "getRecipes/customer (forbidden)",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.a@example.test",
				Roles: []string{},
			}),
			body: &graphql.RawParams{Query: QueryGetRecipes},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": null,
				"errors": [
					{
						"message": "forbidden",
						"path": ["recipes"]
					}
				]
			}
		`},
	},
	{
		name:   "getRecipes/bartender",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			body: &graphql.RawParams{Query: QueryGetRecipes},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": {
					"recipes": [
						{
							"name": "Bombay Sapphire",
							"imageURL": "https://example.test/gintonic.webp",
							"recipes": [
								{
									"name": "Tonic",
									"category": "Cocktail",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Collins", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Bombay Sapphire", "amount": "45ml", "description": null },
										{ "material": null, "amount": null, "description": "Stir" },
										{ "material": "Tonic Water", "amount": "full up", "description": null }
									]
								}
							]
						},
						{
							"name": "Maker's Mark",
							"imageURL": "https://example.test/makersmark.webp",
							"recipes": [
								{
									"name": "Rock",
									"category": "Whisky",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Rock", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Maker's Mark", "amount": "30ml", "description": null }
									]
								},
								{
									"name": "Rock (Double)",
									"category": "Whisky",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Rock", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Maker's Mark", "amount": "60ml", "description": null }
									]
								},
								{
									"name": "Soda",
									"category": "Whisky",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Collins", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Maker's Mark", "amount": "30ml", "description": null },
										{ "material": null, "amount": null, "description": "Stir (Well)" },
										{ "material": "Soda", "amount": "Full up", "description": null }
									]
								},
								{
									"name": "Straight",
									"category": "Whisky",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Straight", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Maker's Mark", "amount": "30ml", "description": null }
									]
								},
								{
									"name": "Water",
									"category": "Whisky",
									"type": { "name": "Build", "description": null },
									"glass": {
										"name": "Straight",
										"imageURL": null,
										"description": null
									},
									"steps": [
										{ "material": "Maker's Mark", "amount": "30ml", "description": "No ice" },
										{ "material": null, "amount": null, "description": "Stir (Well)" },
										{ "material": "Water", "amount": "Full up", "description": null }
									]
								}
							]
						},
						{
							"name": "Phuket Sling",
							"imageURL": "https://example.test/phuketsling.webp",
							"recipes": [
								{
									"name": "Cocktail",
									"category": "Cocktail",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Collins", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Peach Liqueur", "amount": "30ml", "description": null },
										{ "material": "Blue Curacao", "amount": "15ml", "description": null },
										{ "material": "Grapefruit Juice", "amount": "30ml", "description": null },
										{ "material": null, "amount": null, "description": "Stir (Well)" },
										{ "material": "Tonic Water", "amount": "Full up", "description": null }
									]
								},
								{
									"name": "Mocktail",
									"category": "NonAlcholic",
									"type": { "name": "Build", "description": null },
									"glass": { "name": "Collins", "imageURL": null, "description": null },
									"steps": [
										{ "material": "Peach Syrup", "amount": "20ml", "description": null },
										{ "material": "Blue Curacao Syrup", "amount": "10ml", "description": null },
										{ "material": "Grapefruit Juice", "amount": "30ml", "description": null },
										{ "material": null, "amount": null, "description": "Stir (Well)" },
										{ "material": "Tonic Water", "amount": "Full up", "description": null }
									]
								}
							]
						}
					]
				}
			}
		`},
	},
}

func Test_GetRecipe(t *testing.T) {
	for _, tt := range getRecipeTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
