package order

import "fmt"

func (o Order) Validate() error {
	if o.ID == "" {
		return fmt.Errorf("ID cannot be empty")
	}
	if o.CustomerID == "" {
		return fmt.Errorf("CustomerID cannot be empty")
	}
	if o.MenuItemID.ItemName == "" {
		return fmt.Errorf("MenuItemID.ItemName cannot be empty")
	}
	if o.MenuItemID.OptionName == "" {
		return fmt.Errorf("MenuItemID.OptionName cannot be empty")
	}
	for _, timestamp := range o.Timestamps {
		if err := validateStatus(timestamp.Status); err != nil {
			return fmt.Errorf("status \"%s\" is invalid", timestamp.Status)
		}
		if timestamp.Timestamp.IsZero() {
			return fmt.Errorf("timestamp cannot be zero")
		}
	}
	return validateStatus(o.Status)
}

func validateStatus(s Status) error {
	switch s {
	case StatusOrdered:
	case StatusPrepared:
	case StatusDelivered:
	case StatusCanceled:
	default:
		return fmt.Errorf("invalid status")
	}
	return nil
}
