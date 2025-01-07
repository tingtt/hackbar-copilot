package recipe

import "testing"

type ValidateRecipeTypeTest struct {
	RecipeType RecipeType
	Valid      bool
}

var ValidateRecipeTypeTests = []ValidateRecipeTypeTest{
	{RecipeType{Name: "build"}, true},
	{RecipeType{Name: ""}, false},
}

func TestRecipeType_Validate(t *testing.T) {
	t.Parallel()
	for _, tt := range ValidateRecipeTypeTests {
		t.Run("will return nil when valid", func(t *testing.T) {
			t.Parallel()
			err := tt.RecipeType.Validate()
			if tt.Valid && err != nil {
				t.Errorf("RecipeType.Validate() error = %v, want nil", err)
			}
		})
	}
}
