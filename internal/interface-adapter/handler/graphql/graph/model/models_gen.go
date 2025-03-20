// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Cashout struct {
	Checkouts []*Checkout `json:"checkouts"`
	Revenue   float64     `json:"revenue"`
	Timestamp string      `json:"timestamp"`
	StaffID   string      `json:"staffID"`
}

type CashoutInput struct {
	CheckoutIDs []string `json:"checkoutIDs"`
	StaffID     string   `json:"staffID"`
}

type Checkout struct {
	ID            string         `json:"id"`
	CustomerEmail string         `json:"customerEmail"`
	OrderIDs      []string       `json:"orderIDs"`
	Diffs         []*PaymentDiff `json:"diffs"`
	TotalPrice    float64        `json:"totalPrice"`
	PaymentType   CheckoutType   `json:"paymentType"`
	Timestamp     string         `json:"timestamp"`
}

type GlassType struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Description *string `json:"description,omitempty"`
}

type InputAsMenuArgs struct {
	Flavor *string `json:"flavor,omitempty"`
}

type InputAsMenuItemArgs struct {
	ImageURL *string `json:"imageURL,omitempty"`
	Price    float64 `json:"price"`
}

type InputCashoutQuery struct {
	Since string `json:"since"`
	Until string `json:"until"`
}

type InputCheckout struct {
	CustomerEmail string            `json:"customerEmail"`
	OrderIDs      []string          `json:"orderIDs"`
	Diffs         []*InputPriceDiff `json:"diffs"`
	PaymentType   CheckoutType      `json:"paymentType"`
}

type InputGlassType struct {
	Name        string  `json:"name"`
	ImageURL    *string `json:"imageURL,omitempty"`
	Description *string `json:"description,omitempty"`
	Save        *bool   `json:"save,omitempty"`
}

type InputOrder struct {
	MenuItemID   string  `json:"menuItemID"`
	CustomerName *string `json:"customerName,omitempty"`
}

type InputOrderStatusUpdate struct {
	ID     string      `json:"id"`
	Status OrderStatus `json:"status"`
}

type InputPriceDiff struct {
	Price       float64 `json:"price"`
	Description *string `json:"description,omitempty"`
}

type InputRecipe struct {
	Name       string               `json:"name"`
	RecipeType *InputRecipeType     `json:"recipeType,omitempty"`
	GlassType  *InputGlassType      `json:"glassType,omitempty"`
	Steps      []*InputStep         `json:"steps,omitempty"`
	AsMenu     *InputAsMenuItemArgs `json:"asMenu,omitempty"`
}

type InputRecipeGroup struct {
	Name     string           `json:"name"`
	ImageURL *string          `json:"imageURL,omitempty"`
	Recipes  []*InputRecipe   `json:"recipes,omitempty"`
	AsMenu   *InputAsMenuArgs `json:"asMenu,omitempty"`
}

type InputRecipeType struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Save        *bool   `json:"save,omitempty"`
}

type InputStep struct {
	Material    *string `json:"material,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}

type InputStockUpdate struct {
	In  []string `json:"in,omitempty"`
	Out []string `json:"out,omitempty"`
}

type Material struct {
	Name    string `json:"name"`
	InStock bool   `json:"inStock"`
}

type MenuItem struct {
	Name        string            `json:"name"`
	ImageURL    *string           `json:"imageURL,omitempty"`
	Flavor      *string           `json:"flavor,omitempty"`
	Options     []*MenuItemOption `json:"options,omitempty"`
	MinPriceYen float64           `json:"minPriceYen"`
}

type MenuItemOption struct {
	Name       string   `json:"name"`
	ImageURL   *string  `json:"imageURL,omitempty"`
	Materials  []string `json:"materials,omitempty"`
	OutOfStock bool     `json:"outOfStock"`
	PriceYen   float64  `json:"priceYen"`
	Recipe     *Recipe  `json:"recipe,omitempty"`
}

type Mutation struct {
}

type Order struct {
	ID            string                        `json:"id"`
	CustomerEmail string                        `json:"customerEmail"`
	CustomerName  string                        `json:"customerName"`
	MenuItemID    string                        `json:"menuItemID"`
	Timestamps    []*OrderStatusUpdateTimestamp `json:"timestamps"`
	Status        OrderStatus                   `json:"status"`
	Price         float64                       `json:"price"`
}

type OrderStatusUpdateTimestamp struct {
	Status    OrderStatus `json:"status"`
	Timestamp string      `json:"timestamp"`
}

type PaymentDiff struct {
	Price       float64 `json:"price"`
	Description *string `json:"description,omitempty"`
}

type Query struct {
}

type Recipe struct {
	Name  string      `json:"name"`
	Type  *RecipeType `json:"type,omitempty"`
	Glass *GlassType  `json:"glass,omitempty"`
	Steps []*Step     `json:"steps,omitempty"`
}

type RecipeGroup struct {
	Name     string    `json:"name"`
	ImageURL *string   `json:"imageURL,omitempty"`
	Recipes  []*Recipe `json:"recipes,omitempty"`
}

type RecipeType struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type Step struct {
	Material    *string `json:"material,omitempty"`
	Amount      *string `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CheckoutType string

const (
	CheckoutTypeCredit CheckoutType = "CREDIT"
	CheckoutTypeQR     CheckoutType = "QR"
	CheckoutTypeCash   CheckoutType = "CASH"
)

var AllCheckoutType = []CheckoutType{
	CheckoutTypeCredit,
	CheckoutTypeQR,
	CheckoutTypeCash,
}

func (e CheckoutType) IsValid() bool {
	switch e {
	case CheckoutTypeCredit, CheckoutTypeQR, CheckoutTypeCash:
		return true
	}
	return false
}

func (e CheckoutType) String() string {
	return string(e)
}

func (e *CheckoutType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckoutType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckoutType", str)
	}
	return nil
}

func (e CheckoutType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatus string

const (
	OrderStatusOrdered    OrderStatus = "ORDERED"
	OrderStatusPrepared   OrderStatus = "PREPARED"
	OrderStatusDelivered  OrderStatus = "DELIVERED"
	OrderStatusCanceled   OrderStatus = "CANCELED"
	OrderStatusCheckedout OrderStatus = "CHECKEDOUT"
	OrderStatusUnknown    OrderStatus = "UNKNOWN"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusOrdered,
	OrderStatusPrepared,
	OrderStatusDelivered,
	OrderStatusCanceled,
	OrderStatusCheckedout,
	OrderStatusUnknown,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusOrdered, OrderStatusPrepared, OrderStatusDelivered, OrderStatusCanceled, OrderStatusCheckedout, OrderStatusUnknown:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
