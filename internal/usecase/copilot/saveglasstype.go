package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
)

// SaveGlassType implements Copilot.
func (c *copilot) SaveGlassType(gt recipe.GlassType) error {
	err := gt.Validate()
	if err != nil {
		return fmt.Errorf("failed to create glass type: invalid glass type: %w", err)
	}

	err = c.datasource.Recipe().SaveGlassType(gt)
	if err != nil {
		return fmt.Errorf("failed to save glass type: %w", err)
	}
	return nil
}
