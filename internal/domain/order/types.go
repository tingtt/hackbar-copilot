package order

import (
	"time"
)

type Order struct {
	ID         ID
	CustomerID CustomerID
	MenuItemID MenuItemID
	Timestamps []StatusUpdateTimestamp
	Status     Status
	Price      float32
}

type SavedOrder struct {
	Order
	Err error
}

type UUID string

type ID UUID

type CustomerID UUID

type MenuItemID struct {
	GroupName string
	ItemName  string
}

func (m MenuItemID) String() string {
	return m.GroupName + "-" + m.ItemName
}

type Status string

const (
	StatusOrdered    Status = "Ordered"
	StatusPrepared   Status = "Prepared"
	StatusDelivered  Status = "Delivered"
	StatusCanceled   Status = "Canceled"
	StatusCheckedOut Status = "CheckedOut"
)

type StatusUpdateTimestamp struct {
	Status    Status
	Timestamp time.Time
}
