package stocktest

import (
	"hackbar-copilot/internal/domain/stock"
)

var ExampleMaterials = []stock.Material{
	{Name: "Peach liqueur", InStock: true},
	{Name: "Blue curacao", InStock: true},
	{Name: "Grapefruit juice", InStock: true},
	{Name: "Tonic water", InStock: true},
	{Name: "Passoa", InStock: true},
	{Name: "Gin", InStock: true},
	{Name: "Lemon juice", InStock: true},
}

var ExampleMaterialsIter = IterWithNilError(ExampleMaterials)
var ExampleMaterialsMap = func() map[string]bool {
	m := make(map[string]bool)
	for _, v := range ExampleMaterials {
		m[v.Name] = v.InStock
	}
	return m
}()
var ExampleMaterialNames = func() (inStock, outOfStock []string) {
	for _, material := range ExampleMaterials {
		if material.InStock {
			inStock = append(inStock, material.Name)
		} else {
			outOfStock = append(outOfStock, material.Name)
		}
	}
	return inStock, outOfStock
}
