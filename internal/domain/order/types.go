package order

import (
	"time"
)

type Order struct {
	ID
	CustomerID
	MenuItemID
	Timestamps []StatusUpdateTimestamp
	Status
	Price float32
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

type CheckoutType string

const (
	CheckoutTypeCreditCard CheckoutType = "CreditCard"
	CheckoutTypeQR         CheckoutType = "QR"
	CheckoutTypeCash       CheckoutType = "Cash"
)

type StatusUpdateTimestamp struct {
	Status
	Timestamp time.Time
}
