package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_saverLister_SaveRecipeType(t *testing.T) {
	t.Parallel()

	t.Run("may return validation error", func(t *testing.T) {
		t.Parallel()

		argRecipeType := RecipeType{Name: ""}

		s := &saverLister{Repository: nil}
		err := s.SaveRecipeType(argRecipeType)
		assert.Error(t, err)
	})

	t.Run("will call Repository.SaveRecipeType with sanitized recipe type", func(t *testing.T) {
		t.Parallel()

		argRecipeType := RecipeType{
			Name:        "build",
			Description: ptr(""),
		}
		wantArgSaveRecipeType := RecipeType{
			Name:        "build",
			Description: nil,
		}

		mockRepository := new(MockRepository)
		mockRepository.On("SaveRecipeType", mock.Anything).Return(nil)

		s := &saverLister{mockRepository}
		err := s.SaveRecipeType(argRecipeType)
		assert.NoError(t, err)
		mockRepository.AssertCalled(t, "SaveRecipeType", wantArgSaveRecipeType)
	})
}
