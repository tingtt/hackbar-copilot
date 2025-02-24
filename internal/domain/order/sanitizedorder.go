package order

func (o Order) Sanitized() Order {
	sanitized := o
	for i, t := range o.Timestamps {
		sanitized.Timestamps[i] = t.Sanitized()
	}
	return o
}
