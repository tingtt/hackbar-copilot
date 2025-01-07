package recipe

import "fmt"

func (rt *RecipeType) Validate() error {
	if rt.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
