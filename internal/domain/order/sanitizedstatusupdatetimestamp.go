package order

func (t StatusUpdateTimestamp) Sanitized() StatusUpdateTimestamp {
	sanitized := t
	sanitized.Timestamp = t.Timestamp.UTC()
	return sanitized
}
