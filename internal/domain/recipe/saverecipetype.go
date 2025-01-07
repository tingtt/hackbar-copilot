package recipe

import (
	"fmt"
)

// SaveRecipeType implements SaveLister.
func (s *saverLister) SaveRecipeType(rt RecipeType) error {
	if err := rt.Validate(); err != nil {
		return fmt.Errorf("invalid recipe type: %w", err)
	}
	if rt.Description != nil && *rt.Description == "" {
		rt.Description = nil
	}
	return s.Repository.SaveRecipeType(rt)
}
