package menu

import (
	"fmt"
	"slices"
)

func (i ItemOption) Sanitized() ItemOption {
	sanitized := i
	if i.ImageURL != nil && *i.ImageURL == "" {
		sanitized.ImageURL = nil
	}
	slices.Sort(sanitized.Materials)
	sanitized.Materials = slices.Compact(sanitized.Materials)
	return sanitized
}

func (i *ItemOption) Validate() error {
	if i.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
