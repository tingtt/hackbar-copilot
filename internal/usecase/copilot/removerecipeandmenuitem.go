package copilot

import (
	"errors"
	"hackbar-copilot/internal/domain/menu"
)

// RemoveRecipeAndMenuItem implements Copilot.
func (c *copilot) RemoveRecipeAndMenuItem(name string) error {
	err := c.menu.Remove(name)
	if err != nil && !errors.Is(err, menu.ErrNotFound) {
		return err
	}
	return c.recipe.Remove(name)
}
