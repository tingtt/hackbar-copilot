package graph

import (
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"time"
)

type inputCheckout model.InputCheckout

func (i inputCheckout) apply() (order.CustomerEmail, []order.ID, []checkout.Diff, checkout.PaymentType) {
	orderIDs := make([]order.ID, 0, len(i.OrderIDs))
	for _, orderID := range i.OrderIDs {
		orderIDs = append(orderIDs, order.ID(orderID))
	}

	diffs := make([]checkout.Diff, 0, len(i.Diffs))
	for _, diff := range i.Diffs {
		diffs = append(diffs, checkout.Diff{
			Price:       float32(diff.Price),
			Description: diff.Description,
		})
	}

	return order.CustomerEmail(i.CustomerEmail), orderIDs, diffs, checkout.PaymentType(i.PaymentType)
}

type checkout_ checkout.Checkout

func (c checkout_) apply() *model.Checkout {
	m := model.Checkout{
		ID:            string(c.ID),
		CustomerEmail: string(c.CustomerEmail),
		TotalPrice:    float64(c.TotalPrice),
		PaymentType:   model.CheckoutType(c.PaymentType),
		Timestamp:     c.Timestamp.UTC().Format(time.RFC3339),
	}
	for _, orderID := range c.OrderIDs {
		m.OrderIDs = append(m.OrderIDs, string(orderID))
	}
	for _, diff := range c.Diffs {
		m.Diffs = append(m.Diffs, &model.PaymentDiff{
			Price:       float64(diff.Price),
			Description: diff.Description,
		})
	}
	return &m
}

type checkouts_ []checkout.Checkout

func (c checkouts_) apply() []*model.Checkout {
	checkouts := make([]*model.Checkout, 0, len(c))
	for _, checkout := range c {
		checkouts = append(checkouts, checkout_(checkout).apply())
	}
	return checkouts
}
