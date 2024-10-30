package graph

import (
	"context"
	"fmt"
	"net/http"
)

type ContextKey string

const CONTEXT_KEY_HEADER ContextKey = "header"

func loadContext(ctx context.Context) http.Header {
	raw, ok := ctx.Value(CONTEXT_KEY_HEADER).(http.Header)
	if !ok {
		panic(fmt.Sprintf("context[\"%s\"]", CONTEXT_KEY_HEADER))
	}
	return raw
}
