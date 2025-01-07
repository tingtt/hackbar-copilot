package recipe

import (
	"fmt"
)

// Save implements SaveLister.
func (s *saverLister) Save(rg RecipeGroup) error {
	if err := rg.Validate(); err != nil {
		return fmt.Errorf("invalid recipe group: %w", err)
	}
	return s.Repository.Save(rg)
}
