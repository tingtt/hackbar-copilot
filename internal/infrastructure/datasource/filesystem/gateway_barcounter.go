package filesystem

import (
	"hackbar-copilot/internal/usecase/barcounter"
)

var _ barcounter.Gateway = (*barcounterGateway)(nil)

type barcounterGateway struct {
	*gateway
}

// Order implements barcounter.Gateway.
func (g *barcounterGateway) Order() barcounter.OrderSaveListFinder {
	return &g.order
}
