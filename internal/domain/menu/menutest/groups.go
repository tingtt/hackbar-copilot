package menutest

import (
	"hackbar-copilot/internal/domain/menu"
)

func ptr[T any](v T) *T {
	return &v
}

var ExampleGroups = []menu.Group{
	{
		Name:     "Phuket Sling",
		ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
		Flavor:   ptr("Sweet"),
		Items: []menu.Item{
			{
				Name:       "Cocktail",
				ImageURL:   ptr("https://example.com/path/to/image/phuket-sling/cocktail"),
				Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      700,
			},
			{
				Name:       "Mocktail",
				ImageURL:   ptr("https://example.com/path/to/image/phuket-sling/mocktail"),
				Materials:  []string{"Peach syrup", "Blue curacao syrup", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      500,
			},
		},
	},
	{
		Name:     "Passoamoni",
		ImageURL: ptr("https://example.com/path/to/image/passoamoni"),
		Flavor:   ptr("Fruity"),
		Items: []menu.Item{
			{
				Name:       "Cocktail",
				ImageURL:   ptr("https://example.com/path/to/image/passoamoni"),
				Materials:  []string{"Passoa", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      700,
			},
		},
	},
	{
		Name:     "Blue Devil",
		ImageURL: ptr("https://example.com/path/to/image/blue-devil"),
		Flavor:   ptr("Medium sweet and dry"),
		Items: []menu.Item{
			{
				Name:       "Cocktail",
				ImageURL:   ptr("https://example.com/path/to/image/blue-devil"),
				Materials:  []string{"Gin", "Blue curacao", "Lemon juice"},
				OutOfStock: false,
				Price:      700,
			},
		},
	},
}

var ExampleGroupsIter = IterWithNilError(ExampleGroups)
