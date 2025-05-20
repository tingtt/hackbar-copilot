package test

import (
	"hackbar-copilot/internal/utils/sliceutil"

	"github.com/99designs/gqlgen/graphql"
)

func InitialOrders(token string) []IntegrationTestRequest {
	return sliceutil.Map(PramsOrder(), func(p graphql.RawParams) IntegrationTestRequest {
		return IntegrationTestRequest{
			token: token,
			body:  &p,
		}
	})
}

func PramsOrder() []graphql.RawParams {
	params := []graphql.RawParams{}
	for _, variable := range VariablesOrder {
		params = append(params, graphql.RawParams{
			Query:     QueryOrder,
			Variables: map[string]any{"input": variable},
		})
	}
	return params
}

var VariablesOrder = []map[string]any{
	{
		"menuItemName":       "Maker's Mark",
		"menuItemOptionName": "Rock",
		"customerName":       "Customer A",
	},
	{
		"menuItemName":       "Maker's Mark",
		"menuItemOptionName": "Rock",
		"customerName":       "Customer A",
	},
	{
		"menuItemName":       "Phuket Sling",
		"menuItemOptionName": "Mocktail",
		"customerName":       "Customer A",
	},
}
