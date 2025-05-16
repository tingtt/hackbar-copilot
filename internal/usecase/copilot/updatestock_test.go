package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/domain/stock/stocktest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateStockTest struct {
	Name string

	// args
	InStockMaterials    []string
	OutOfStockMaterials []string

	// current
	Menu []menu.Item

	// behavior
	UpdatedMaterials []stock.Material

	// expect
	SaveMenuExpectCalls []menu.Item
}

var updateStockTests = []UpdateStockTest{
	{
		Name:                "change to out of stock",
		InStockMaterials:    []string{},
		OutOfStockMaterials: []string{"Peach liqueur"},
		Menu: []menu.Item{
			{
				Name: "Phuket Sling",
				Options: []menu.ItemOption{
					{
						Name:       "Cocktail",
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						Price:      700,
					},
				},
			},
		},
		UpdatedMaterials: []stock.Material{
			{Name: "Peach liqueur", InStock: false},
			{Name: "Blue curacao", InStock: true},
			{Name: "Grapefruit juice", InStock: true},
			{Name: "Tonic water", InStock: true},
		},
		SaveMenuExpectCalls: []menu.Item{
			{
				Name: "Phuket Sling",
				Options: []menu.ItemOption{
					{
						Name:       "Cocktail",
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: true,
						Price:      700,
					},
				},
			},
		},
	},
	{
		Name:                "change to in stock",
		InStockMaterials:    []string{"Peach liqueur"},
		OutOfStockMaterials: []string{},
		Menu: []menu.Item{
			{
				Name: "Phuket Sling",
				Options: []menu.ItemOption{
					{
						Name:       "Cocktail",
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: true,
						Price:      700,
					},
				},
			},
		},
		UpdatedMaterials: []stock.Material{
			{Name: "Peach liqueur", InStock: true},
			{Name: "Blue curacao", InStock: true},
			{Name: "Grapefruit juice", InStock: true},
			{Name: "Tonic water", InStock: true},
		},
		SaveMenuExpectCalls: []menu.Item{
			{
				Name: "Phuket Sling",
				Options: []menu.ItemOption{
					{
						Name:       "Cocktail",
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice", "Tonic water"},
						OutOfStock: false,
						Price:      700,
					},
				},
			},
		},
	},
}

func Test_copilot_UpdateStock(t *testing.T) {
	t.Parallel()

	t.Run("may call menu.Save with menu stock status changed", func(t *testing.T) {
		for _, tt := range updateStockTests {
			t.Run(tt.Name, func(t *testing.T) {
				t.Parallel()
				stockMock := new(MockStock)
				stockMock.On("All").Return(stocktest.IterWithNilError(tt.UpdatedMaterials))
				stockMock.On("Save", mock.Anything, mock.Anything).Return(nil)
				menuMock := new(MockMenu)
				menuMock.On("All").Return(menutest.IterWithNilError(tt.Menu))
				menuMock.On("Save", mock.Anything).Return(nil)
				gateway := MockGateway{menu: menuMock, stock: stockMock}

				c := &copilot{&gateway}
				err := c.UpdateStock(tt.InStockMaterials, tt.OutOfStockMaterials)

				assert.NoError(t, err)
				for _, saveMenuExpectCall := range tt.SaveMenuExpectCalls {
					menuMock.AssertCalled(t, "Save", saveMenuExpectCall)
				}
			})
		}
	})

	t.Run("will call stock.SaveLister.Save", func(t *testing.T) {
		for _, tt := range updateStockTests {
			t.Run(tt.Name, func(t *testing.T) {
				t.Parallel()
				stockMock := new(MockStock)
				stockMock.On("All").Return(stocktest.IterWithNilError(tt.UpdatedMaterials))
				stockMock.On("Save", mock.Anything, mock.Anything).Return(nil)
				menuMock := new(MockMenu)
				menuMock.On("All").Return(menutest.IterWithNilError(tt.Menu))
				menuMock.On("Save", mock.Anything).Return(nil)
				gateway := MockGateway{menu: menuMock, stock: stockMock}

				c := &copilot{&gateway}
				err := c.UpdateStock(tt.InStockMaterials, tt.OutOfStockMaterials)

				assert.NoError(t, err)
				stockMock.AssertCalled(t, "Save", tt.InStockMaterials, tt.OutOfStockMaterials)
			})
		}
	})
}
