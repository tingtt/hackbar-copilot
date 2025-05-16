package filesystem

import "hackbar-copilot/internal/usecase/cashier"

var _ cashier.Gateway = (*cashierGateway)(nil)

type cashierGateway struct {
	*gateway
}

// Order implements cashier.Gateway.
func (g *cashierGateway) Order() cashier.OrderListRemover {
	return &g.order
}

// Checkout implements cashier.Gateway.
func (g *cashierGateway) Checkout() cashier.CheckoutSaveListRemover {
	return &g.checkout
}

// Cashout implements cashier.Gateway.
func (g *cashierGateway) Cashout() cashier.CashoutSaveLister {
	return &g.cashout
}
