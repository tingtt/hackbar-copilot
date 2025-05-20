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

var orderListTests = []IntegrationTest{
	{
		name: "getUncheckedOrdersCustomer/list customer A's orders",
		before: append(append(
			InitialSaveRecipes(),
			InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.a@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
			}), "Customer A")...),
			InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.b@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer B"},
			}), "Customer B")...,
		),
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
				assert.NotNil(t, res.Data.UncheckedOrders, msgAndArgs...)
				assert.Contains(t,
					sliceutil.Map(res.Data.UncheckedOrders, func(o model.Order) string {
						return o.CustomerName
					}),
					"Customer A",
					msgAndArgs...,
				)
				assert.NotContains(t,
					sliceutil.Map(res.Data.UncheckedOrders, func(o model.Order) string {
						return o.CustomerName
					}),
					"Customer B",
					msgAndArgs...,
				)
			},
		},
	},
	{
		name: "getUncheckedOrdersCustomer/list all unchecked orders",
		before: append(append(
			InitialSaveRecipes(),
			InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.a@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
			}), "Customer A")...),
			InitialOrders(graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.b@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer B"},
			}), "Customer B")...,
		),
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "bartender@example.test",
				Roles:  []string{"bartender"},
				GitHub: &jwtclaims.ClaimsGitHub{ID: "bartender"},
			}),
			body: &graphql.RawParams{
				Query: QueryGetUncheckedOrders,
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
				assert.NotNil(t, res.Data.UncheckedOrders, msgAndArgs...)
				orderCustomerNames := sliceutil.Map(res.Data.UncheckedOrders, func(o model.Order) string { return o.CustomerName })
				assert.Contains(t, orderCustomerNames, "Customer A", msgAndArgs...)
				assert.Contains(t, orderCustomerNames, "Customer B", msgAndArgs...)
				orderCustomerEmails := sliceutil.Map(res.Data.UncheckedOrders, func(o model.Order) string { return o.CustomerEmail })
				assert.Contains(t, orderCustomerEmails, "customer.a@example.test", msgAndArgs...)
				assert.Contains(t, orderCustomerEmails, "customer.b@example.test", msgAndArgs...)
			},
		},
	},
}

func Test_ListOrder(t *testing.T) {
	for _, tt := range orderListTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
