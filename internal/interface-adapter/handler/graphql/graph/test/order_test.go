package test

import (
	"context"
	"encoding/json"
	"fmt"
	graphqlhandler "hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var orderTests = []IntegrationTest{
	{
		name: "order/customer A",
		before: append(
			InitialSaveRecipes(),
			IntegrationTestRequest{
				// register customer data
				token: graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.a@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
				}),
				body: &graphql.RawParams{Query: QueryGetUserInfo},
			},
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.a@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
			}),
			body: &graphql.RawParams{
				Query: QueryOrder,
				Variables: map[string]any{
					"input": map[string]any{
						"menuItemName":       "Maker's Mark",
						"menuItemOptionName": "Rock",
						"customerName":       "Customer A",
					},
				},
			},
		},
		want: IntegrationTestWantResponse{assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
			var res Response[struct {
				Order model.Order `json:"order"`
			}]
			if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
				assert.Fail(t,
					fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
					msgAndArgs...,
				)
				return
			}
			assert.Nil(t, res.Errors, msgAndArgs...)
			assert.Equal(t, res.Data.Order.MenuID.ItemName, "Maker's Mark", msgAndArgs...)
			assert.Equal(t, res.Data.Order.MenuID.OptionName, "Rock", msgAndArgs...)
			assert.Equal(t, res.Data.Order.CustomerEmail, "customer.a@example.test", msgAndArgs...)
			assert.Equal(t, res.Data.Order.CustomerName, "Customer A", msgAndArgs...)
			assert.Equal(t, res.Data.Order.Status, model.OrderStatusOrdered, msgAndArgs...)
			assert.Equal(t, res.Data.Order.Price, float64(800), msgAndArgs...)
		}},
		after: []IntegrationTest{
			{
				name: "getUncheckedOrders/customer A orders",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{
						Email:  "customer.a@example.test",
						Roles:  []string{},
						Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
					}),
					body: &graphql.RawParams{
						Query: QueryGetUncheckedOrdersCustomer,
					},
				},
				want: IntegrationTestWantResponse{
					assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
						var res Response[struct {
							UncheckedOrders []model.Order `json:"uncheckedOrdersCustomer"`
						}]
						if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
							assert.Fail(t,
								fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
								msgAndArgs...,
							)
							return
						}
						assert.Nil(t, res.Errors, msgAndArgs...)
						assert.Equal(t, len(res.Data.UncheckedOrders), 1, msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].MenuID.ItemName, "Maker's Mark", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].MenuID.OptionName, "Rock", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].CustomerEmail, "customer.a@example.test", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].CustomerName, "Customer A", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].Status, model.OrderStatusOrdered, msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].Price, float64(800), msgAndArgs...)
					},
				},
			},
			{
				name: "getUserInfo/customer name confirmed",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{
						Email:  "customer.a@example.test",
						Roles:  []string{},
						Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
					}),
					body: &graphql.RawParams{
						Query: QueryGetUserInfo,
					},
				},
				want: IntegrationTestWantResponse{
					bodyJSON: `
						{
							"data": {
								"userInfo": {
									"email": "customer.a@example.test",
									"name": "Customer A",
									"nameConfirmed": true
								}
							}
						}
					`,
				},
			},
		},
	},
	{
		name:   "order/bartender as customer",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			body: &graphql.RawParams{
				Query: QueryOrder,
				Variables: map[string]any{
					"input": map[string]any{
						"menuItemName":       "Maker's Mark",
						"menuItemOptionName": "Rock",
						"customerName":       "Customer A",
						"customerEmail":      "customer.a@example.test",
					},
				},
			},
		},
		want: IntegrationTestWantResponse{assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
			var res Response[struct {
				Order model.Order `json:"order"`
			}]
			if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
				assert.Fail(t,
					fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
					msgAndArgs...,
				)
				return
			}
			assert.Nil(t, res.Errors, msgAndArgs...)
			assert.Equal(t, res.Data.Order.MenuID.ItemName, "Maker's Mark", msgAndArgs...)
			assert.Equal(t, res.Data.Order.MenuID.OptionName, "Rock", msgAndArgs...)
			assert.Equal(t, res.Data.Order.CustomerEmail, "customer.a@example.test", msgAndArgs...)
			assert.Equal(t, res.Data.Order.CustomerName, "Customer A", msgAndArgs...)
			assert.Equal(t, res.Data.Order.Status, model.OrderStatusOrdered, msgAndArgs...)
			assert.Equal(t, res.Data.Order.Price, float64(800), msgAndArgs...)
		}},
		after: []IntegrationTest{
			{
				name: "getUncheckedOrders/customer A orders",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{
						Email:  "customer.a@example.test",
						Roles:  []string{},
						Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
					}),
					body: &graphql.RawParams{
						Query: QueryGetUncheckedOrdersCustomer,
					},
				},
				want: IntegrationTestWantResponse{
					assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
						var res Response[struct {
							UncheckedOrders []model.Order `json:"uncheckedOrdersCustomer"`
						}]
						if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
							assert.Fail(t,
								fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
								msgAndArgs...,
							)
							return
						}
						assert.Nil(t, res.Errors, msgAndArgs...)
						assert.Equal(t, len(res.Data.UncheckedOrders), 1, msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].MenuID.ItemName, "Maker's Mark", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].MenuID.OptionName, "Rock", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].CustomerEmail, "customer.a@example.test", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].CustomerName, "Customer A", msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].Status, model.OrderStatusOrdered, msgAndArgs...)
						assert.Equal(t, res.Data.UncheckedOrders[0].Price, float64(800), msgAndArgs...)
					},
				},
			},
		},
	},
	{
		name:   "order/bartender as customer (error cause customerName not specified)",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			body: &graphql.RawParams{
				Query: QueryOrder,
				Variables: map[string]any{
					"input": map[string]any{
						"menuItemName":       "Maker's Mark",
						"menuItemOptionName": "Rock",
						"customerEmail":      "customer.a@example.test",
					},
				},
			},
		},
		want: IntegrationTestWantResponse{
			bodyJSON: `
				{
					"errors": [
						{
							"message": "'customerName' cannot be empty for ordering as another account",
							"path": ["order"]
						}
					],
					"data": null
				}
			`,
		},
	},
	{
		name:   "order/customer as another customer (forbidden)",
		before: InitialSaveRecipes(),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.b@example.test",
			}),
			body: &graphql.RawParams{
				Query: QueryOrder,
				Variables: map[string]any{
					"input": map[string]any{
						"menuItemName":       "Maker's Mark",
						"menuItemOptionName": "Rock",
						"customerEmail":      "customer.a@example.test",
						"customerName":       "Customer A",
					},
				},
			},
		},
		want: IntegrationTestWantResponse{
			bodyJSON: `
				{
					"errors": [
						{
							"message": "forbidden",
							"path": ["order"]
						}
					],
					"data": null
				}
			`,
		},
	},
}

func Test_Order(t *testing.T) {
	for _, tt := range orderTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
