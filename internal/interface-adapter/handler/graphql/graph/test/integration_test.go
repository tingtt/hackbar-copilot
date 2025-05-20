package test

import (
	"context"
	"encoding/json"
	"fmt"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/test/graphqltest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
)

type IntegrationTest struct {
	name    string
	before  []IntegrationTestRequest
	request IntegrationTestRequest
	want    IntegrationTestWantResponse
	after   []IntegrationTest
}

type IntegrationTestRequest struct {
	token string

	// body for the GraphQL query
	//   Not used if resolveContext is specified
	body *graphql.RawParams

	// createBody is used to resolve the context for the GraphQL query
	createBody func(ctx context.Context) *graphql.RawParams
}

type ContextKey string

type IntegrationTestWantResponse struct {
	bodyJSON string
	assert   func(t *testing.T, ctx context.Context, got *httptest.ResponseRecorder, msgAndArgs ...any)
}

type Response[T any] struct {
	Data   T `json:"data"`
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors"`
}

func run(t *testing.T, handler http.Handler, ctx context.Context, tt IntegrationTest, path string) {
	for i, before := range tt.before {
		gotBefore := httptest.NewRecorder()
		if before.createBody != nil {
			before.body = before.createBody(ctx)
		}
		handler.ServeHTTP(gotBefore, graphqltest.Request(before.body, http.Header{
			"Authorization": []string{"Bearer " + before.token},
		}))
		var res Response[any]
		if err := json.Unmarshal(gotBefore.Body.Bytes(), &res); err != nil {
			assert.Fail(t,
				fmt.Sprintf(
					"Before '%s' failed with errors: Input ('%s') needs to be valid json.\nJSON parsing error: '%s'",
					tt.name, gotBefore.Body.String(), err.Error(),
				),
				path,
			)
			return
		}
		assert.Empty(t, res.Errors, fmt.Sprintf("Before '%s' failed with errors: %v", tt.name, res.Errors))
		ctx = context.WithValue(ctx, ContextKey(fmt.Sprintf("%s.before.%d", path, i)), res.Data)
	}

	got := httptest.NewRecorder()

	if tt.request.createBody != nil {
		tt.request.body = tt.request.createBody(ctx)
	}
	handler.ServeHTTP(got, graphqltest.Request(tt.request.body, http.Header{
		"Authorization": []string{"Bearer " + tt.request.token},
	}))

	if tt.want.bodyJSON != "" {
		assert.JSONEq(t, tt.want.bodyJSON, got.Body.String(), path)
	}
	if tt.want.assert != nil {
		tt.want.assert(t, ctx, got, path)
	}

	var res Response[any]
	if err := json.Unmarshal(got.Body.Bytes(), &res); err != nil {
		assert.Fail(t,
			fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", got.Body.String(), err.Error()), path,
		)
		return
	}
	ctx = context.WithValue(ctx, ContextKey(fmt.Sprintf("%s.run", path)), res.Data)

	for i, after := range tt.after {
		run(t, handler, ctx, after, fmt.Sprintf("%s.after[%d]", path, i))
	}
}
