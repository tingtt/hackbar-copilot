package copilot

import (
	"hackbar-copilot/internal/domain/recipe"
)

// SaveGlassType implements Copilot.
func (c *copilot) SaveGlassType(gt recipe.GlassType) error {
	return c.recipe.SaveGlassType(gt)
}
