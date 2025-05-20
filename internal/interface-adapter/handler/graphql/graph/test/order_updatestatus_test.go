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

var orderUpdateTests = []IntegrationTest{
	{
		name: "updateOrderStatus/customer A (forbidden)",
		before: append(InitialSaveRecipes(), InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
			Email:  "customer.a@example.test",
			Roles:  []string{},
			Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
		}))...),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.a@example.test",
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				return &graphql.RawParams{
					Query: QueryUpdateOrderStatus,
					Variables: map[string]any{
						"input": map[string]any{
							"id":     ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"],
							"status": "CANCELED",
						},
					},
				}
			},
		},
		want: IntegrationTestWantResponse{
			bodyJSON: `
				{
					"data": null,
					"errors": [
						{
							"message": "forbidden",
							"path": ["updateOrderStatus"]
						}
					]
				}
			`,
		},
		after: []IntegrationTest{},
	},
	{
		name: "updateOrderStatus/delivered",
		before: append(InitialSaveRecipes(), InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
			Email:  "customer.a@example.test",
			Roles:  []string{},
			Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
		}))...),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "bartender@example.test",
				Roles:  []string{"bartender"},
				GitHub: &jwtclaims.ClaimsGitHub{ID: "bartender"},
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				return &graphql.RawParams{
					Query: QueryUpdateOrderStatus,
					Variables: map[string]any{
						"input": map[string]any{
							"id":     ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"],
							"status": "CANCELED",
						},
					},
				}
			},
		},
		want: IntegrationTestWantResponse{
			assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
				var res Response[struct {
					UpdateOrderStatus model.Order `json:"updateOrderStatus"`
				}]
				if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
					assert.Fail(t,
						fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
						msgAndArgs...,
					)
					return
				}
				assert.Nil(t, res.Errors, msgAndArgs...)
				assert.Equal(t, res.Data.UpdateOrderStatus.ID, ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"], msgAndArgs...)
				assert.Equal(t, res.Data.UpdateOrderStatus.Status, model.OrderStatusCanceled, msgAndArgs...)
			},
		},
		after: []IntegrationTest{
			{
				name: "getUncheckedOrders/customer A orders (updated)",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{Email: "customer.a@example.test"}),
					body: &graphql.RawParams{
						Query: QueryGetUncheckedOrder,
					},
				},
				want: IntegrationTestWantResponse{
					assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
						var res Response[struct {
							UncheckedOrders []model.Order `json:"uncheckedOrders"`
						}]
						if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
							assert.Fail(t,
								fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
								msgAndArgs...,
							)
							return
						}
						assert.Nil(t, res.Errors, msgAndArgs...)
						for _, order := range res.Data.UncheckedOrders {
							if order.ID == ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"] {
								assert.Equal(t, order.Status, model.OrderStatusCanceled, msgAndArgs...)
							} else {
								assert.Equal(t, order.Status, model.OrderStatusOrdered, msgAndArgs...)
							}
						}
					},
				},
			},
		},
	},
}

func Test_OrderUpdate(t *testing.T) {
	for _, tt := range orderUpdateTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
