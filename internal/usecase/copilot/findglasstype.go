package copilot

import (
	"fmt"
	"hackbar-copilot/internal/domain/recipe"
)

// FindGlassType implements Copilot.
func (c *copilot) FindGlassType() (map[string]recipe.GlassType, error) {
	glassTypes := map[string]recipe.GlassType{}
	for gt, err := range c.datasource.Recipe().AllGlassTypes() {
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve glass types: %w", err)
		}
		glassTypes[gt.Name] = gt
	}
	return glassTypes, nil
}
