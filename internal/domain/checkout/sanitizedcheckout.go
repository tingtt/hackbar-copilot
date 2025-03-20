package checkout

func (c *Checkout) Sanitized() Checkout {
	return Checkout{
		ID:            c.ID,
		CustomerEmail: c.CustomerEmail,
		OrderIDs:      c.OrderIDs,
		Diffs:         c.Diffs,
		TotalPrice:    c.TotalPrice,
		PaymentType:   c.PaymentType,
		Timestamp:     c.Timestamp.UTC(),
	}
}
