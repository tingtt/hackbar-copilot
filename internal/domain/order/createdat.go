package order

import (
	"time"
)

func (o *Order) CreatedAt() time.Time {
	for _, t := range o.Timestamps {
		if t.Status == StatusPrepared {
			return t.Timestamp
		}
	}
	panic("order does not have a created timestamp")
}
