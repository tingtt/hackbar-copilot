package copilot

import "hackbar-copilot/internal/domain/recipe"

// FindGlassType implements Copilot.
func (c *copilot) FindGlassType() (map[string]recipe.GlassType, error) {
	glassTypes := map[string]recipe.GlassType{}
	for gt, err := range c.recipe.AllGlassTypes() {
		if err != nil {
			return nil, err
		}
		glassTypes[gt.Name] = gt
	}
	return glassTypes, nil
}
