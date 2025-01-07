package recipe

import (
	"fmt"
)

func (r *Recipe) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if r.Type == "" {
		return fmt.Errorf("type cannot be empty")
	}
	if r.Glass == "" {
		return fmt.Errorf("glass cannot be empty")
	}
	for i, s := range r.Steps {
		if err := s.Validate(); err != nil {
			return fmt.Errorf("step %d is invalid: %w", i+1, err)
		}
	}
	return nil
}
