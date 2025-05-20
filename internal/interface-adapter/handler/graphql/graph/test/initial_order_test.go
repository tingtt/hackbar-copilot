package test

import (
	"hackbar-copilot/internal/utils/sliceutil"

	"maps"

	"github.com/99designs/gqlgen/graphql"
)

func InitialOrders(token, customerName string) []IntegrationTestRequest {
	return sliceutil.Map(PramsOrder(customerName), func(p graphql.RawParams) IntegrationTestRequest {
		return IntegrationTestRequest{
			token: token,
			body:  &p,
		}
	})
}

func PramsOrder(customerName string) []graphql.RawParams {
	params := []graphql.RawParams{}
	for _, variable := range VariablesOrder {
		copiedVariable := make(map[string]any, len(variable))
		maps.Copy(copiedVariable, variable)

		copiedVariable["customerName"] = customerName
		params = append(params, graphql.RawParams{
			Query:     QueryOrder,
			Variables: map[string]any{"input": copiedVariable},
		})
	}
	return params
}

var VariablesOrder = []map[string]any{
	{
		"menuItemName":       "Maker's Mark",
		"menuItemOptionName": "Rock",
	},
	{
		"menuItemName":       "Maker's Mark",
		"menuItemOptionName": "Rock",
	},
	{
		"menuItemName":       "Phuket Sling",
		"menuItemOptionName": "Mocktail",
	},
}
