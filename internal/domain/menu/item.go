package menu

import (
	"fmt"
)

func (g Item) Sanitized() Item {
	sanitized := g
	sanitized.Options = make([]ItemOption, 0, len(g.Options))

	if g.ImageURL != nil && *g.ImageURL == "" {
		sanitized.ImageURL = nil
	}
	if g.Flavor != nil && *g.Flavor == "" {
		sanitized.Flavor = nil
	}
	for _, item := range g.Options {
		sanitized.Options = append(sanitized.Options, item.Sanitized())
	}
	return sanitized
}

func (g *Item) Validate() error {
	if g.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	for _, i := range g.Options {
		if err := i.Validate(); err != nil {
			return fmt.Errorf("item \"%s\" is invalid: %w", i.Name, err)
		}
	}
	return nil
}
