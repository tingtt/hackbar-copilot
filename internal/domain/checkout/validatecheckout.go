package checkout

import (
	"errors"
	"fmt"
)

func (c *Checkout) Validate() error {
	if c.ID == "" {
		return errors.New("ID cannot be empty")
	}
	if c.CustomerEmail == "" {
		return errors.New("CustomerEmail cannot be empty")
	}
	if c.TotalPrice <= 0 {
		return errors.New("total price cannot be less than or equal to zero")
	}
	if c.Timestamp.IsZero() {
		return fmt.Errorf("timestamp cannot be zero")
	}
	return validatePaymentType(c.PaymentType)
}

func validatePaymentType(t PaymentType) error {
	switch t {
	case CheckoutTypeCreditCard:
	case CheckoutTypeQR:
	case CheckoutTypeCash:
	default:
		return fmt.Errorf("invalid payment type")
	}
	return nil
}
