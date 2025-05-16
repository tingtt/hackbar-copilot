package order

import "time"

func (o Order) ApplyStatus(status Status, timestamp time.Time) Order {
	o.Status = status
	o.Timestamps = append(o.Timestamps, StatusUpdateTimestamp{
		Status:    status,
		Timestamp: timestamp,
	})
	return o
}
