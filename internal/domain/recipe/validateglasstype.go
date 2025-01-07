package recipe

import "fmt"

func (gt *GlassType) Validate() error {
	if gt.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
