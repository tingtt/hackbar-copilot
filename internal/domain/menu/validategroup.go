package menu

import (
	"fmt"
)

func (g *Group) Validate() error {
	if g.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	for _, i := range g.Items {
		if err := i.Validate(); err != nil {
			return fmt.Errorf("item \"%s\" is invalid: %w", i.Name, err)
		}
	}
	return nil
}
