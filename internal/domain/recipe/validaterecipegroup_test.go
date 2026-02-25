package recipe

import (
	"testing"
)

type ValidateRecipeGroupTest struct {
	RecipeGroup RecipeGroup
	Valid       bool
}

var validateRecipeGroupTests = []ValidateRecipeGroupTest{
	{
		RecipeGroup: RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: new("https://example.com/path/to/image"),
			Recipes: []Recipe{
				{
					Name:     "Cocktail",
					Category: "Cocktail",
					Type:     "build",
					Glass:    "collins",
					Steps: []Step{
						{
							Material: new("Peach Liqueur"),
							Amount:   new("30ml"),
						},
						{
							Description: new("Stir"),
						},
					},
				},
			},
		},
		Valid: true,
	},
	{
		RecipeGroup: RecipeGroup{
			Name: "",
		},
		Valid: false,
	},
	{
		RecipeGroup: RecipeGroup{
			Name: "Phuket Sling",
			Recipes: []Recipe{
				{
					Name:     "",
					Category: "",
					Type:     "",
					Glass:    "",
					Steps:    []Step{},
				},
			},
		},
		Valid: false,
	},
}

func TestRecipeGroup_Validate(t *testing.T) {
	t.Parallel()

	for _, tt := range validateRecipeGroupTests {
		t.Run("will return nil when valid", func(t *testing.T) {
			t.Parallel()
			err := tt.RecipeGroup.Validate()
			if tt.Valid && err != nil {
				t.Errorf("RecipeGroup.Validate() error = %v, want nil", err)
			}
		})
	}
}
