package test

import (
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"testing"

	graphqlhandler "hackbar-copilot/internal/interface-adapter/handler/graphql"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var getMenuTests = []IntegrationTest{
	{
		name:   "getMenu/customer (no recipe)",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.a@example.test",
				Roles: []string{},
			}),
			body: &graphql.RawParams{Query: QueryGetMenu},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": {
					"menu": [
						{
							"name": "Maker's Mark",
							"imageURL": "https://example.test/makersmark.webp",
							"flavor": null,
							"options": [
								{
									"name": "Rock",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": null
								},
								{
									"name": "Rock (Double)",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 1600,
									"recipe": null
								},
								{
									"name": "Soda",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark", "Soda"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": null
								},
								{
									"name": "Straight",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": null
								}
							],
							"minPriceYen": 800
						},
						{
							"name": "Phuket Sling",
							"imageURL": "https://example.test/phuketsling.webp",
							"flavor": "Sweet and Sour",
							"options": [
								{
									"name": "Cocktail",
									"category": "Cocktail",
									"imageURL": null,
									"materials": [ "Peach Liqueur", "Blue Curacao", "Grapefruit Juice", "Tonic Water" ],
									"outOfStock": false,
									"priceYen": 700,
									"recipe": null
								},
								{
									"name": "Mocktail",
									"category": "NonAlcholic",
									"imageURL": null,
									"materials": [ "Peach Syrup", "Blue Curacao Syrup", "Grapefruit Juice", "Tonic Water" ],
									"outOfStock": false,
									"priceYen": 700,
									"recipe": null
								}
							],
							"minPriceYen": 700
						}
					]
				}
			}
		`},
	},
	{
		name:   "getMenu/bartender (with recipe)",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			body: &graphql.RawParams{Query: QueryGetMenu},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": {
					"menu": [
						{
							"name": "Maker's Mark",
							"imageURL": "https://example.test/makersmark.webp",
							"flavor": null,
							"options": [
								{
									"name": "Rock",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": {
										"name": "Rock",
										"category": "Whisky",
										"type": { "name": "Build", "description": null },
										"glass": { "name": "Rock", "imageURL": null, "description": null },
										"steps": [
											{ "material": "Maker's Mark", "amount": "30ml", "description": null }
										]
									}
								},
								{
									"name": "Rock (Double)",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 1600,
									"recipe": {
										"name": "Rock (Double)",
										"category": "Whisky",
										"type": { "name": "Build", "description": null },
										"glass": { "name": "Rock", "imageURL": null, "description": null },
										"steps": [
											{ "material": "Maker's Mark", "amount": "60ml", "description": null }
										]
									}
								},
								{
									"name": "Soda",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark", "Soda"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": {
										"name": "Soda",
										"category": "Whisky",
										"type": { "name": "Build", "description": null },
										"glass": { "name": "Collins", "imageURL": null, "description": null },
										"steps": [
											{ "material": "Maker's Mark", "amount": "30ml", "description": null },
											{ "material": null, "amount": null, "description": "Stir (Well)" },
											{ "material": "Soda", "amount": "Full up", "description": null }
										]
									}
								},
								{
									"name": "Straight",
									"category": "Whisky",
									"imageURL": null,
									"materials": ["Maker's Mark"],
									"outOfStock": false,
									"priceYen": 800,
									"recipe": {
										"name": "Straight",
										"category": "Whisky",
										"type": { "name": "Build", "description": null },
										"glass": { "name": "Straight", "imageURL": null, "description": null },
										"steps": [
											{ "material": "Maker's Mark", "amount": "30ml", "description": null }
										]
									}
								}
							],
							"minPriceYen": 800
						},
						{
							"name": "Phuket Sling",
							"imageURL": "https://example.test/phuketsling.webp",
							"flavor": "Sweet and Sour",
							"options": [
								{
									"name": "Cocktail",
									"category": "Cocktail",
									"imageURL": null,
									"materials": [ "Peach Liqueur", "Blue Curacao", "Grapefruit Juice", "Tonic Water" ],
									"outOfStock": false,
									"priceYen": 700,
									"recipe": {
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
									}
								},
								{
									"name": "Mocktail",
									"category": "NonAlcholic",
									"imageURL": null,
									"materials": [ "Peach Syrup", "Blue Curacao Syrup", "Grapefruit Juice", "Tonic Water" ],
									"outOfStock": false,
									"priceYen": 700,
									"recipe": {
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
								}
							],
							"minPriceYen": 700
						}
					]
				}
			}
		`},
	},
}

func Test_GetMenu(t *testing.T) {
	for _, tt := range getMenuTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t, graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())), tt)
		})
	}
}
