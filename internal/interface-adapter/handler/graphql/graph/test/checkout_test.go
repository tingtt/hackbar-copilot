package test

import (
	"context"
	"encoding/json"
	"fmt"
	graphqlhandler "hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"hackbar-copilot/internal/utils/sliceutil"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var checkoutTests = []IntegrationTest{
	{
		name: "checkout/customer A orders",
		before: append(
			InitialSaveRecipes(),
			InitialOrders(
				graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.a@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
				}), "Customer A",
			)...,
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				orderIDs := []any{
					ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.4")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.5")).(map[string]any)["order"].(map[string]any)["id"],
				}
				return &graphql.RawParams{
					Query: QueryCheckout,
					Variables: map[string]any{
						"input": map[string]any{
							"customerEmail": "customer.a@example.test",
							"orderIDs":      orderIDs,
							"diffs": []map[string]any{
								{
									"price":       1000,
									"description": "charge",
								},
							},
							"paymentType": "CREDIT",
						},
					},
				}
			},
		},
		want: IntegrationTestWantResponse{assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
			var res Response[struct {
				Checkout model.Checkout `json:"checkout"`
			}]
			if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
				assert.Fail(t,
					fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
					msgAndArgs...,
				)
				return
			}

			assert.Nil(t, res.Errors, msgAndArgs...)
			assert.Equal(t, res.Data.Checkout.CustomerEmail, "customer.a@example.test")
			{
				wantOrderIDs := []any{
					ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.4")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.5")).(map[string]any)["order"].(map[string]any)["id"],
				}
				gotOrderIDs := sliceutil.Map(res.Data.Checkout.Orders, func(o *model.Order) string { return o.ID })
				assert.ElementsMatch(t, wantOrderIDs, gotOrderIDs, msgAndArgs...)
			}
			assert.False(t, sliceutil.Some(res.Data.Checkout.Orders, func(o *model.Order) bool {
				return o.Status != model.OrderStatusCheckedout
			}), msgAndArgs...)
			assert.NotNil(t, res.Data.Checkout.Diffs, msgAndArgs...)
			assert.Equal(t, len(res.Data.Checkout.Diffs), 1, msgAndArgs...)
			assert.Equal(t, res.Data.Checkout.Diffs[0].Price, float64(1000), msgAndArgs...)
			assert.Equal(t, res.Data.Checkout.TotalPrice, float64(3300), msgAndArgs...)
			assert.Equal(t, res.Data.Checkout.PaymentType, model.CheckoutTypeCredit, msgAndArgs...)
		}},
		after: []IntegrationTest{
			{
				name: "getUncheckedOrders/customer A orders cleared",
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
					bodyJSON: `
						{
							"data": {
								"uncheckedOrdersCustomer": []
							}
						}
					`,
				},
			},
			{
				name: "getCheckouts",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{
						Email: "bartender@example.test",
						Roles: []string{"bartender"},
					}),
					body: &graphql.RawParams{
						Query: QueryGetUncashedCheckouts,
					},
				},
				want: IntegrationTestWantResponse{
					assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
						var res Response[struct {
							UncashedoutCheckouts []model.Checkout `json:"uncashedoutCheckouts"`
						}]
						if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
							assert.Fail(t,
								fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
								msgAndArgs...,
							)
							return
						}

						assert.Nil(t, res.Errors, msgAndArgs...)
						assert.Equal(t, len(res.Data.UncashedoutCheckouts), 1, msgAndArgs...)
						wantCheckoutID := ctx.Value(ContextKey(".run")).(map[string]any)["checkout"].(map[string]any)["id"]
						assert.Equal(t, res.Data.UncashedoutCheckouts[0].ID, wantCheckoutID, msgAndArgs...)
					},
				},
			},
		},
	},
	{
		name: "checkout/customer (forbidden)",
		before: append(
			InitialSaveRecipes(),
			InitialOrders(
				graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.a@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
				}), "Customer A",
			)...,
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.a@example.test",
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				orderIDs := []any{
					ctx.Value(ContextKey(".before.3")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.4")).(map[string]any)["order"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.5")).(map[string]any)["order"].(map[string]any)["id"],
				}
				return &graphql.RawParams{
					Query: QueryCheckout,
					Variables: map[string]any{
						"input": map[string]any{
							"customerEmail": "customer.a@example.test",
							"orderIDs":      orderIDs,
							"diffs":         []any{},
							"paymentType":   "CREDIT",
						},
					},
				}
			},
		},
		want: IntegrationTestWantResponse{
			bodyJSON: `
				{
					"errors": [
						{
							"message": "forbidden",
							"path": ["checkout"]
						}
					],
					"data": null
				}
			`,
		},
	},
}

func Test_Checkout(t *testing.T) {
	for _, tt := range checkoutTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
