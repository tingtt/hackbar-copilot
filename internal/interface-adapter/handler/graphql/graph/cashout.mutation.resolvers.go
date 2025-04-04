package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"errors"
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/domain/checkout"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// Cashout is the resolver for the cashout field.
func (r *mutationResolver) Cashout(ctx context.Context, input model.CashoutInput) (*model.Cashout, error) {
	email, err := r.authAdapter.GetEmail(ctx)
	if /* unauthorized */ err != nil {
		return nil, err
	}
	if !r.authAdapter.HasBartenderRole(ctx) {
		return nil, errors.New("forbidden")
	}

	checkoutIDs := make([]checkout.ID, 0, len(input.CheckoutIDs))
	for _, checkoutID := range input.CheckoutIDs {
		checkoutIDs = append(checkoutIDs, checkout.ID(checkoutID))
	}

	cashout, err := r.Cashier.Cashout(cashout.StaffID(email), checkoutIDs)
	if err != nil {
		return nil, err
	}
	return cashout_(cashout).apply(), nil
}
