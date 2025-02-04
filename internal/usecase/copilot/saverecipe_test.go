package copilot

import (
	"hackbar-copilot/internal/domain/menu"
	"hackbar-copilot/internal/domain/menu/menutest"
	"hackbar-copilot/internal/domain/recipe"
	"hackbar-copilot/internal/domain/recipe/recipetest"
	"hackbar-copilot/internal/domain/stock"
	"hackbar-copilot/internal/domain/stock/stocktest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SaveRecipeTest struct {
	Name string

	// args
	RecipeGroup recipe.RecipeGroup

	// current
	Menu      []menu.Group
	Materials []stock.Material

	// expect
	SaveAsMenuGroupExpectCall *SaveRecipeTestExpectCall
}

type SaveRecipeTestExpectCall struct {
	SaveMenuGroup    menu.Group
	SaveNewMaterials []string
}

var saveRecipeTests = []SaveRecipeTest{
	{
		Name: "save recipe not in menu",
		RecipeGroup: recipe.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
			Recipes: []recipe.Recipe{
				{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []recipe.Step{
						{
							Material: ptr("Peach liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Material: ptr("Blue curacao"),
							Amount:   ptr("15ml"),
						},
						{
							Material: ptr("Grapefruit juice"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
						},
						{
							Material: ptr("Tonic water"),
							Amount:   ptr("Full up"),
						},
					},
				},
			},
		},
		Menu:      []menu.Group{},
		Materials: []stock.Material{},
	},
	{
		Name: "save recipe used in menu",
		RecipeGroup: recipe.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
			Recipes: []recipe.Recipe{
				{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []recipe.Step{
						{
							Material: ptr("Peach liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Material: ptr("Blue curacao"),
							Amount:   ptr("15ml"),
						},
						{
							Material: ptr("Grapefruit juice"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
						},
						{
							Material: ptr("Tonic water"),
							Amount:   ptr("Full up"),
						},
					},
				},
			},
		},
		Menu: []menu.Group{
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
				},
			},
		},
		Materials: []stock.Material{
			{Name: "Peach liqueur", InStock: true},
			{Name: "Blue curacao", InStock: true},
			{Name: "Grapefruit juice", InStock: true},
			{Name: "Tonic water", InStock: true},
		},
		SaveAsMenuGroupExpectCall: &SaveRecipeTestExpectCall{
			SaveMenuGroup: menu.Group{
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
				},
			},
			SaveNewMaterials: []string{},
		},
	},
	{
		Name: "save recipe used in menu with new material",
		RecipeGroup: recipe.RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
			Recipes: []recipe.Recipe{
				{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []recipe.Step{
						{
							Material: ptr("Peach liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Material: ptr("Blue curacao"),
							Amount:   ptr("15ml"),
						},
						{
							Material: ptr("Grapefruit juice"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
						},
						{
							Material: ptr("Tonic water"),
							Amount:   ptr("Full up"),
						},
					},
				},
			},
		},
		Menu: []menu.Group{
			{
				Name:     "Phuket Sling",
				ImageURL: ptr("https://example.com/path/to/image/phuket-sling"),
				Flavor:   ptr("Sweet"),
				Items: []menu.Item{
					{
						Name:       "Cocktail",
						ImageURL:   ptr("https://example.com/path/to/image/phuket-sling/cocktail"),
						Materials:  []string{"Peach liqueur", "Blue curacao", "Grapefruit juice"},
						OutOfStock: false,
						Price:      700,
					},
				},
			},
		},
		Materials: []stock.Material{
			{Name: "Peach liqueur", InStock: true},
			{Name: "Blue curacao", InStock: true},
			{Name: "Grapefruit juice", InStock: true},
		},
		SaveAsMenuGroupExpectCall: &SaveRecipeTestExpectCall{
			SaveMenuGroup: menu.Group{
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
				},
			},
			SaveNewMaterials: []string{"Tonic water"},
		},
	},
}

func Test_copilot_SaveRecipe(t *testing.T) {
	t.Parallel()

	t.Run("may call SaveAsMenuGroup, when related menu group found", func(t *testing.T) {
		t.Parallel()
		for _, tt := range saveRecipeTests {
			t.Run(tt.Name, func(t *testing.T) {
				t.Parallel()
				recipeMock := new(MockRecipeSaveLister)
				recipeMock.On("All").Return(recipetest.IterWithNilError([]recipe.RecipeGroup{tt.RecipeGroup}))
				recipeMock.On("Save", mock.Anything).Return(nil)
				stockMock := new(MockStockSaveLister)
				stockMock.On("All").Return(stocktest.IterWithNilError(tt.Materials))
				stockMock.On("Save", mock.Anything, mock.Anything).Return(nil)
				menuMock := new(MockMenuSaveLister)
				menuMock.On("All").Return(menutest.IterWithNilError(tt.Menu))
				menuMock.On("Save", mock.Anything).Return(nil)

				c := &copilot{recipe: recipeMock, menu: menuMock, stock: stockMock}
				err := c.SaveRecipe(tt.RecipeGroup)

				assert.NoError(t, err)
				if tt.SaveAsMenuGroupExpectCall != nil {
					menuMock.AssertCalled(t, "Save", tt.SaveAsMenuGroupExpectCall.SaveMenuGroup)
					if len(tt.SaveAsMenuGroupExpectCall.SaveNewMaterials) > 0 {
						stockMock.AssertCalled(t, "Save", tt.SaveAsMenuGroupExpectCall.SaveNewMaterials, []string(nil))
					} else {
						stockMock.AssertNotCalled(t, "Save")
					}
				} else {
					menuMock.AssertNotCalled(t, "Save")
					stockMock.AssertNotCalled(t, "Save")
				}
			})
		}
	})

	t.Run("will call recipe.SaveLister.Save", func(t *testing.T) {
		t.Parallel()
		for _, tt := range saveRecipeTests {
			t.Run(tt.Name, func(t *testing.T) {
				t.Parallel()
				recipeMock := new(MockRecipeSaveLister)
				recipeMock.On("All").Return(recipetest.IterWithNilError([]recipe.RecipeGroup{tt.RecipeGroup}))
				recipeMock.On("Save", mock.Anything).Return(nil)
				stockMock := new(MockStockSaveLister)
				stockMock.On("All").Return(stocktest.IterWithNilError(tt.Materials))
				stockMock.On("Save", mock.Anything, mock.Anything).Return(nil)
				menuMock := new(MockMenuSaveLister)
				menuMock.On("All").Return(menutest.IterWithNilError(tt.Menu))
				menuMock.On("Save", mock.Anything).Return(nil)

				c := &copilot{recipe: recipeMock, menu: menuMock, stock: stockMock}
				err := c.SaveRecipe(tt.RecipeGroup)

				assert.NoError(t, err)
				recipeMock.AssertCalled(t, "Save", tt.RecipeGroup)
			})
		}
	})
}
