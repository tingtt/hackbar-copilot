package order

import (
	"hackbar-copilot/internal/domain/user"
	"time"
)

type Order struct {
	ID            ID
	CustomerEmail CustomerEmail
	CustomerName  string
	MenuItemID    MenuItemID
	Timestamps    []StatusUpdateTimestamp
	Status        Status
	Price         float32
}

type SavedOrder struct {
	Order
	Err error
}

type UUID string

type ID UUID

type CustomerEmail user.Email

type MenuItemID struct {
	ItemName   string
	OptionName string
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
