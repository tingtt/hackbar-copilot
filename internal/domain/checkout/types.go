package checkout

import (
	"hackbar-copilot/internal/domain/order"
	"time"
)

type Checkout struct {
	ID          ID
	CustomerID  order.CustomerID
	OrderIDs    []order.ID
	Diffs       []Diff
	TotalPrice  float32
	PaymentType PaymentType
	Timestamp   time.Time
}

type UUID string

type ID UUID

type Diff struct {
	Price       float32
	Description *string
}

type PaymentType string

const (
	CheckoutTypeCreditCard PaymentType = "CreditCard"
	CheckoutTypeQR         PaymentType = "QR"
	CheckoutTypeCash       PaymentType = "Cash"
)
