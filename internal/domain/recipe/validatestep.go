package recipe

import (
	"fmt"
)

func (s *Step) Validate() error {
	if s.Material == nil {
		if s.Amount != nil {
			return fmt.Errorf("amount cannot be set without material")
		}
		if s.Description == nil {
			return fmt.Errorf("description cannot be empty, if material is not set")
		}
	}
	return nil
}
