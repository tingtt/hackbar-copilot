package graph

import (
	"fmt"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/domain/order"
	orderadapter "hackbar-copilot/internal/interface-adapter/handler/graphql/adapter/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"time"
)

type inputCheckout model.InputCheckout

func (i inputCheckout) apply() (order.CustomerEmail, []order.ID, []checkout.Diff, checkout.PaymentType, error) {
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

	paymentType, err := inputCheckoutType(i.PaymentType).apply()
	if err != nil {
		return "", nil, nil, "", err
	}

	return order.CustomerEmail(i.CustomerEmail), orderIDs, diffs, paymentType, nil
}

type inputCheckoutType model.CheckoutType

func (i inputCheckoutType) apply() (checkout.PaymentType, error) {
	switch model.CheckoutType(i) {
	case model.CheckoutTypeCredit:
		return checkout.CheckoutTypeCreditCard, nil
	case model.CheckoutTypeQR:
		return checkout.CheckoutTypeQR, nil
	case model.CheckoutTypeCash:
		return checkout.CheckoutTypeCash, nil
	default:
		return "", fmt.Errorf("invalid checkout type")
	}
}

type paymentType checkout.PaymentType

func (p paymentType) apply() model.CheckoutType {
	switch checkout.PaymentType(p) {
	case checkout.CheckoutTypeCreditCard:
		return model.CheckoutTypeCredit
	case checkout.CheckoutTypeQR:
		return model.CheckoutTypeQR
	case checkout.CheckoutTypeCash:
		return model.CheckoutTypeCash
	default:
		return "unknown"
	}
}

type checkout_ checkout.Checkout

func (c checkout_) apply() *model.Checkout {
	orderAdapter := orderadapter.NewOutputAdapter()

	m := model.Checkout{
		ID:            string(c.ID),
		CustomerEmail: string(c.CustomerEmail),
		TotalPrice:    float64(c.TotalPrice),
		PaymentType:   paymentType(c.PaymentType).apply(),
		Timestamp:     c.Timestamp.UTC().Format(time.RFC3339),
	}
	m.Orders = orderAdapter.Orders(c.Orders)
	m.Diffs = diffs_(c.Diffs).apply()
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

type diffs_ []checkout.Diff

func (d diffs_) apply() []*model.PaymentDiff {
	diffs := make([]*model.PaymentDiff, 0, len(d))
	for _, diff := range d {
		diffs = append(diffs, diff_(diff).apply())
	}
	return diffs
}

type diff_ checkout.Diff

func (d diff_) apply() *model.PaymentDiff {
	return &model.PaymentDiff{
		Price:       float64(d.Price),
		Description: d.Description,
	}
}
