package menu

import (
	"slices"
)

func (i Item) Sanitized() Item {
	sanitized := i
	if i.ImageURL != nil && *i.ImageURL == "" {
		sanitized.ImageURL = nil
	}
	slices.Sort(sanitized.Materials)
	sanitized.Materials = slices.Compact(sanitized.Materials)
	return sanitized
}
