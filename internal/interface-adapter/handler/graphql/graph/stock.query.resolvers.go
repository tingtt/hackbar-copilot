package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"hackbar-copilot/internal/usecase/copilot"
)

// Materials is the resolver for the materials field.
func (r *queryResolver) Materials(ctx context.Context) ([]*model.Material, error) {
	materials, err := r.Copilot.Materials(copilot.SortMaterialByName())
	if err != nil {
		return nil, err
	}

	var result []*model.Material
	for _, material := range materials {
		result = append(result, &model.Material{
			Name:    material.Name,
			InStock: material.InStock,
		})
	}
	return result, nil
}
