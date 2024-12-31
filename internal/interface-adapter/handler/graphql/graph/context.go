package graph

import (
	"context"
	"fmt"
	"net/http"
)

type ContextKey string

const ContextKeyHeader ContextKey = "header"

//lint:ignore U1000 we will use when implementing authentication
func loadContext(ctx context.Context) http.Header {
	raw, ok := ctx.Value(ContextKeyHeader).(http.Header)
	if !ok {
		panic(fmt.Sprintf("context[\"%s\"]", ContextKeyHeader))
	}
	return raw
}
