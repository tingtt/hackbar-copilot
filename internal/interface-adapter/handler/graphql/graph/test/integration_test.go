package test

import (
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
	body  *graphql.RawParams
}

type IntegrationTestWantResponse struct {
	bodyJSON string
	assert   func(t *testing.T, got *httptest.ResponseRecorder, msgAndArgs ...any)
}

type Response[T any] struct {
	Data   T `json:"data"`
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors"`
}

func run(t *testing.T, handler http.Handler, tt IntegrationTest, msgAndArgs ...any) {
	for _, before := range tt.before {
		gotBefore := httptest.NewRecorder()
		handler.ServeHTTP(gotBefore, graphqltest.Request(before.body, http.Header{
			"Authorization": []string{"Bearer " + before.token},
		}))
		var res Response[any]
		if err := json.Unmarshal(gotBefore.Body.Bytes(), &res); err != nil {
			assert.Fail(t,
				fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", gotBefore.Body.String(), err.Error()),
				msgAndArgs...,
			)
			return
		}
		assert.Empty(t, res.Errors, fmt.Sprintf("Before '%s' failed with errors: %v", tt.name, res.Errors))
	}

	got := httptest.NewRecorder()

	handler.ServeHTTP(got, graphqltest.Request(tt.request.body, http.Header{
		"Authorization": []string{"Bearer " + tt.request.token},
	}))

	if tt.want.bodyJSON != "" {
		assert.JSONEq(t, tt.want.bodyJSON, got.Body.String(), msgAndArgs...)
	}
	if tt.want.assert != nil {
		tt.want.assert(t, got, msgAndArgs...)
	}
	for _, after := range tt.after {
		run(t, handler, after, "Failed assert after '%s'", tt.name+".after[]."+after.name)
	}
}
