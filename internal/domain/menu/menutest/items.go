package menutest

import (
	"hackbar-copilot/internal/domain/menu"
)

var ExampleItems = []menu.Item{
	{
		Name:     "Phuket Sling",
		ImageURL: new("https://example.com/path/to/image/phuket-sling"),
		Flavor:   new("Sweet"),
		Options: []menu.ItemOption{
			{
				Name:       "Cocktail",
				Category:   "Cocktail",
				ImageURL:   new("https://example.com/path/to/image/phuket-sling/cocktail"),
				Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      700,
			},
			{
				Name:       "Mocktail",
				Category:   "Mocktail",
				ImageURL:   new("https://example.com/path/to/image/phuket-sling/mocktail"),
				Materials:  []string{"Peach syrup", "Blue curacao syrup", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      500,
			},
		},
	},
	{
		Name:     "Passoamoni",
		ImageURL: new("https://example.com/path/to/image/passoamoni"),
		Flavor:   new("Fruity"),
		Options: []menu.ItemOption{
			{
				Name:       "Cocktail",
				Category:   "Cocktail",
				ImageURL:   new("https://example.com/path/to/image/passoamoni"),
				Materials:  []string{"Passoa", "Grapefruit juice", "Tonic water"},
				OutOfStock: false,
				Price:      700,
			},
		},
	},
	{
		Name:     "Blue Devil",
		ImageURL: new("https://example.com/path/to/image/blue-devil"),
		Flavor:   new("Medium sweet and dry"),
		Options: []menu.ItemOption{
			{
				Name:       "Cocktail",
				Category:   "Cocktail",
				ImageURL:   new("https://example.com/path/to/image/blue-devil"),
				Materials:  []string{"Gin", "Blue curacao", "Lemon juice"},
				OutOfStock: false,
				Price:      700,
			},
		},
	},
}

var ExampleItemsIter = IterWithNilError(ExampleItems)
