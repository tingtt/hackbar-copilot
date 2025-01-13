package menu

import (
	"fmt"
)

func (i *Item) Validate() error {
	if i.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
