package recipe

import (
	"fmt"
)

// SaveGlassType implements SaveLister.
func (s *saverLister) SaveGlassType(gt GlassType) error {
	if err := gt.Validate(); err != nil {
		return fmt.Errorf("invalid glass type: %w", err)
	}
	if gt.ImageURL != nil && *gt.ImageURL == "" {
		gt.ImageURL = nil
	}
	if gt.Description != nil && *gt.Description == "" {
		gt.Description = nil
	}
	return s.Repository.SaveGlassType(gt)
}
