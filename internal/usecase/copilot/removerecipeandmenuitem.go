package copilot

import (
	"errors"
	"fmt"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
)

// RemoveRecipeAndMenuItem implements Copilot.
func (c *copilot) RemoveRecipeAndMenuItem(name string) error {
	err := c.datasource.Menu().Remove(name)
	if err != nil && !errors.Is(err, usecaseutils.ErrNotFound) {
		return fmt.Errorf("failed to remove menu item: %w", err)
	}
	err = c.datasource.Recipe().Remove(name)
	if err != nil && !errors.Is(err, usecaseutils.ErrNotFound) {
		return fmt.Errorf("failed to remove recipe: %w", err)
	}
	return nil
}
