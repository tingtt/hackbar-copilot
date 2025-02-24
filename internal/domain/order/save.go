package order

import (
	"fmt"
)

func (s *saveListListener) Save(o Order) error {
	sanitizedOrder := o.Sanitized()
	if err := sanitizedOrder.Validate(); err != nil {
		return fmt.Errorf("invalid order: %w", err)
	}
	return s.Repository.Save(sanitizedOrder)
}
