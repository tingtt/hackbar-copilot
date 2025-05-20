package graphqltest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/99designs/gqlgen/graphql"
)

func Request(body *graphql.RawParams, header http.Header) *http.Request {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	for key := range header {
		req.Header.Set(key, header.Get(key))
	}

	return req
}
