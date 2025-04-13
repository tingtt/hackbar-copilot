package recipe

import (
	"fmt"
)

func (rg *RecipeGroup) Validate() error {
	if rg.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if len(rg.Recipes) == 0 {
		return fmt.Errorf("recipes cannot be empty")
	}
	for _, r := range rg.Recipes {
		if err := r.Validate(); err != nil {
			return fmt.Errorf("recipe \"%s\" is invalid: %w", r.Name, err)
		}
	}
	return nil
}
