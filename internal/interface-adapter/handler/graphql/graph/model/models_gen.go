// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type SaveRecipeResult interface {
	IsSaveRecipeResult()
}

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

type InputAsMenuItemArgs struct {
	Flavor *string `json:"flavor,omitempty"`
	Remove *bool   `json:"remove,omitempty"`
}

type InputAsMenuItemOptionArgs struct {
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
	MenuItemName       string  `json:"menuItemName"`
	MenuItemOptionName string  `json:"menuItemOptionName"`
	CustomerName       *string `json:"customerName,omitempty"`
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
	Name       string                     `json:"name"`
	Category   string                     `json:"category"`
	RecipeType *InputRecipeType           `json:"recipeType,omitempty"`
	GlassType  *InputGlassType            `json:"glassType,omitempty"`
	Steps      []*InputStep               `json:"steps"`
	Remove     *bool                      `json:"remove,omitempty"`
	AsMenu     *InputAsMenuItemOptionArgs `json:"asMenu,omitempty"`
}

type InputRecipeGroup struct {
	Name     string               `json:"name"`
	ImageURL *string              `json:"imageURL,omitempty"`
	Replace  *bool                `json:"replace,omitempty"`
	Recipes  []*InputRecipe       `json:"recipes"`
	Remove   *bool                `json:"remove,omitempty"`
	AsMenu   *InputAsMenuItemArgs `json:"asMenu,omitempty"`
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

type MenuID struct {
	ItemName   string `json:"itemName"`
	OptionName string `json:"optionName"`
}

type MenuItem struct {
	Name        string            `json:"name"`
	ImageURL    *string           `json:"imageURL,omitempty"`
	Flavor      *string           `json:"flavor,omitempty"`
	Options     []*MenuItemOption `json:"options"`
	MinPriceYen float64           `json:"minPriceYen"`
}

type MenuItemOption struct {
	Name       string   `json:"name"`
	Category   string   `json:"category"`
	ImageURL   *string  `json:"imageURL,omitempty"`
	Materials  []string `json:"materials"`
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
	MenuID        *MenuID                       `json:"menuID"`
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
	Name     string      `json:"name"`
	Category string      `json:"category"`
	Type     *RecipeType `json:"type,omitempty"`
	Glass    *GlassType  `json:"glass,omitempty"`
	Steps    []*Step     `json:"steps"`
}

type RecipeGroup struct {
	Name     string    `json:"name"`
	ImageURL *string   `json:"imageURL,omitempty"`
	Recipes  []*Recipe `json:"recipes"`
}

func (RecipeGroup) IsSaveRecipeResult() {}

type RecipeType struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type RemovedRecipeGroup struct {
	Name string `json:"name"`
}

func (RemovedRecipeGroup) IsSaveRecipeResult() {}

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
