package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"errors"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// UpdateStock is the resolver for the updateStock field.
func (r *mutationResolver) UpdateStock(ctx context.Context, input model.InputStockUpdate) ([]*model.Material, error) {
	_, err := r.authAdapter.GetEmail(ctx)
	if /* unauthorized */ err != nil {
		return nil, err
	}
	if !r.authAdapter.HasBartenderRole(ctx) {
		return nil, errors.New("forbidden")
	}

	err = r.Copilot.UpdateStock(input.In, input.Out)
	if err != nil {
		return nil, err
	}
	return r.Query().Materials(ctx)
}
