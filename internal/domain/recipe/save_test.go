package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_saverLister_Save(t *testing.T) {
	t.Parallel()

	t.Run("may return validation error", func(t *testing.T) {
		t.Parallel()

		for _, tt := range validateRecipeGroupTests {
			if !tt.Valid {
				s := &saverLister{Repository: nil}
				err := s.Save(tt.RecipeGroup)
				assert.Error(t, err)
			}
		}
	})

	t.Run("will call Repository.Save", func(t *testing.T) {
		t.Parallel()

		arg := RecipeGroup{
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
		}

		mockRepository := new(MockRepository)
		mockRepository.On("Save", mock.Anything).Return(nil)

		s := &saverLister{mockRepository}
		err := s.Save(arg)
		assert.NoError(t, err)
		mockRepository.AssertCalled(t, "Save", arg)
	})
}
