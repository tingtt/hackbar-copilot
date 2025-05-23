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
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var cashoutTests = []IntegrationTest{
	{
		name: "cashout",
		before: append(append(append(append(
			InitialSaveRecipes(),
			InitialOrders(
				graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.a@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
				}), "Customer A",
			)...),
			IntegrationTestRequest{
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
								"diffs": []any{
									map[string]any{
										"price":       1000,
										"description": "charge",
									},
								},
								"paymentType": "CREDIT",
							},
						},
					}
				},
			}),
			InitialOrders(
				graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.b@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer B"},
				}), "Customer B",
			)...),
			IntegrationTestRequest{
				token: graphqltest.NewToken(jwtclaims.Claims{
					Email: "bartender@example.test",
					Roles: []string{"bartender"},
				}),
				createBody: func(ctx context.Context) *graphql.RawParams {
					orderIDs := []any{
						ctx.Value(ContextKey(".before.7")).(map[string]any)["order"].(map[string]any)["id"],
						ctx.Value(ContextKey(".before.8")).(map[string]any)["order"].(map[string]any)["id"],
						ctx.Value(ContextKey(".before.9")).(map[string]any)["order"].(map[string]any)["id"],
					}
					return &graphql.RawParams{
						Query: QueryCheckout,
						Variables: map[string]any{
							"input": map[string]any{
								"customerEmail": "customer.b@example.test",
								"orderIDs":      orderIDs,
								"diffs": []any{
									map[string]any{
										"price":       1000,
										"description": "charge",
									},
								},
								"paymentType": "QR",
							},
						},
					}
				},
			},
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "bartender@example.test",
				Roles: []string{"bartender"},
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				checkoutIDs := []any{
					ctx.Value(ContextKey(".before.6")).(map[string]any)["checkout"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.10")).(map[string]any)["checkout"].(map[string]any)["id"],
				}
				return &graphql.RawParams{
					Query: QueryCashout,
					Variables: map[string]any{
						"input": map[string]any{
							"checkoutIDs": checkoutIDs,
							"staffID":     "bartender@example.test",
						},
					},
				}
			},
		},
		want: IntegrationTestWantResponse{assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
			var res Response[struct {
				Cashout model.Cashout `json:"cashout"`
			}]
			if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
				assert.Fail(t,
					fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
					msgAndArgs...,
				)
				return
			}

			assert.Nil(t, res.Errors, msgAndArgs...)
			assert.NotNil(t, res.Data.Cashout.Checkouts, msgAndArgs...)
			assert.Equal(t, 2, len(res.Data.Cashout.Checkouts), msgAndArgs...)
			{
				wantCheckoutIDs := []any{
					ctx.Value(ContextKey(".before.6")).(map[string]any)["checkout"].(map[string]any)["id"],
					ctx.Value(ContextKey(".before.10")).(map[string]any)["checkout"].(map[string]any)["id"],
				}
				gotCheckoutIDs := sliceutil.Map(res.Data.Cashout.Checkouts, func(c *model.Checkout) string { return c.ID })
				assert.ElementsMatch(t, wantCheckoutIDs, gotCheckoutIDs, msgAndArgs...)
			}
			assert.Equal(t, float64(6600), res.Data.Cashout.Revenue, msgAndArgs...)
			assert.Equal(t, "bartender@example.test", res.Data.Cashout.StaffID)
		}},
		after: []IntegrationTest{
			{
				name: "getCheckouts/cleared",
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
					bodyJSON: `
						{
							"data": {
								"uncashedoutCheckouts": []
							}
						}
					`,
				},
			},
			{
				name: "getCashouts",
				request: IntegrationTestRequest{
					token: graphqltest.NewToken(jwtclaims.Claims{
						Email: "bartender@example.test",
						Roles: []string{"bartender"},
					}),
					body: &graphql.RawParams{
						Query: QueryGetCashouts,
						Variables: map[string]any{
							"input": map[string]any{
								"since": time.Now().Format(time.RFC3339),
								"until": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
							},
						},
					},
				},
				want: IntegrationTestWantResponse{
					assert: func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any) {
						var res Response[struct {
							Cashouts []model.Cashout `json:"cashouts"`
						}]
						if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
							assert.Fail(t,
								fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()),
								msgAndArgs...,
							)
							return
						}

						assert.Nil(t, res.Errors, msgAndArgs...)
						assert.Equal(t, len(res.Data.Cashouts), 1, msgAndArgs...)
						wantCashoutRevenue := ctx.Value(ContextKey(".run")).(map[string]any)["cashout"].(map[string]any)["revenue"]
						assert.Equal(t, res.Data.Cashouts[0].Revenue, wantCashoutRevenue, msgAndArgs...)
					},
				},
			},
		},
	},
	{
		name: "cashout (forbidden)",
		before: append(append(
			InitialSaveRecipes(),
			InitialOrders(
				graphqltest.NewToken(jwtclaims.Claims{
					Email:  "customer.a@example.test",
					Roles:  []string{},
					Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
				}), "Customer A",
			)...),
			IntegrationTestRequest{
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
								"diffs":         []any{},
								"paymentType":   "CREDIT",
							},
						},
					}
				},
			},
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email: "customer.a@example.test",
			}),
			createBody: func(ctx context.Context) *graphql.RawParams {
				checkoutIDs := []any{
					ctx.Value(ContextKey(".before.6")).(map[string]any)["checkout"].(map[string]any)["id"],
				}
				return &graphql.RawParams{
					Query: QueryCashout,
					Variables: map[string]any{
						"input": map[string]any{
							"checkoutIDs": checkoutIDs,
							"staffID":     "customer.a@example.test",
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
							"path": ["cashout"]
						}
					],
					"data": null
				}
			`,
		},
	},
}

func Test_Cashout(t *testing.T) {
	for _, tt := range cashoutTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
