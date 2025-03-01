package checkout

import "fmt"

func (s *saveLister) Save(c Checkout) error {
	sanitizedCheckout := c.Sanitized()
	if err := sanitizedCheckout.Validate(); err != nil {
		return fmt.Errorf("invalid checkout: %w", err)
	}
	return s.Repository.Save(sanitizedCheckout)
}
