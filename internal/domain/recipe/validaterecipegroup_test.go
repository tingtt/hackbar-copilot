package recipe

import (
	"testing"
)

func ptr(s string) *string {
	return &s
}

type ValidateRecipeGroupTest struct {
	RecipeGroup RecipeGroup
	Valid       bool
}

var validateRecipeGroupTests = []ValidateRecipeGroupTest{
	{
		RecipeGroup: RecipeGroup{
			Name:     "Phuket Sling",
			ImageURL: ptr("https://example.com/path/to/image"),
			Recipes: []Recipe{
				{
					Name:  "Cocktail",
					Type:  "build",
					Glass: "collins",
					Steps: []Step{
						{
							Material: ptr("Peach Liqueur"),
							Amount:   ptr("30ml"),
						},
						{
							Description: ptr("Stir"),
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
					Name:  "",
					Type:  "",
					Glass: "",
					Steps: []Step{},
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
