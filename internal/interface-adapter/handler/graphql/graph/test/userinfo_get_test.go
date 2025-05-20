package test

import (
	"context"
	graphqlhandler "hackbar-copilot/internal/interface-adapter/handler/graphql"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tingtt/oauth2rbac/pkg/jwtclaims"
)

var getUserInfoTests = []IntegrationTest{
	{
		name: "getUserInfo/customer",
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "customer.a@example.test",
				Roles:  []string{},
				Google: &jwtclaims.ClaimsGoogle{Username: "Customer A"},
			}),
			body: &graphql.RawParams{Query: QueryGetUserInfo},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": {
					"userInfo": {
						"email": "customer.a@example.test",
						"name": "Customer A",
						"nameConfirmed": false
					}
				}
			}
		`},
	},
	{
		name: "getUserInfo/bartender",
		request: IntegrationTestRequest{
			token: graphqltest.NewToken(jwtclaims.Claims{
				Email:  "bartender@example.test",
				Roles:  []string{"bartender"},
				GitHub: &jwtclaims.ClaimsGitHub{ID: "bartender"},
			}),
			body: &graphql.RawParams{Query: QueryGetUserInfo},
		},
		want: IntegrationTestWantResponse{bodyJSON: `
			{
				"data": {
					"userInfo": {
						"email": "bartender@example.test",
						"name": "bartender",
						"nameConfirmed": false
					}
				}
			}
		`},
	},
}

func Test_GetUserInfo(t *testing.T) {
	for _, tt := range getUserInfoTests {
		t.Run(tt.name, func(t *testing.T) {
			run(t,
				graphqlhandler.NewHandler(graphqltest.Dependencies(t.TempDir())),
				context.Background(), tt, "",
			)
		})
	}
}
