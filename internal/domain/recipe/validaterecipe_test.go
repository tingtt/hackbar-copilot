package recipe

import "testing"

type ValidateRecipeTest struct {
	Recipe Recipe
	Valid  bool
}

var validateRecipeTests = []ValidateRecipeTest{
	{
		Recipe: Recipe{
			Name:     "Cocktail",
			Category: "Cocktail",
			Type:     "build",
			Glass:    "collins",
			Steps: []Step{
				{
					Material: ptr("Peach Liqueur"),
					Amount:   ptr("30ml"),
				},
			},
		},
		Valid: true,
	},
	{
		Recipe: Recipe{
			Name:  "",
			Type:  "build",
			Glass: "collins",
			Steps: []Step{},
		},
		Valid: false,
	},
	{
		Recipe: Recipe{
			Name:  "Cocktail",
			Type:  "",
			Glass: "collins",
			Steps: []Step{},
		},
		Valid: false,
	},
	{
		Recipe: Recipe{
			Name:  "Cocktail",
			Type:  "build",
			Glass: "",
			Steps: []Step{},
		},
		Valid: false,
	},
	{
		Recipe: Recipe{
			Name:  "Cocktail",
			Type:  "build",
			Glass: "collins",
			Steps: []Step{
				{ /** invalid step */ },
			},
		},
		Valid: false,
	},
}

func TestRecipe_Validate(t *testing.T) {
	t.Parallel()

	for _, tt := range validateRecipeTests {
		t.Run("will return nil when valid", func(t *testing.T) {
			t.Parallel()
			err := tt.Recipe.Validate()
			if tt.Valid && err != nil {
				t.Errorf("Recipe.Validate() error = %v, want nil", err)
			}
		})
	}
}
