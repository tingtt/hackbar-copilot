package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Email CustomerEmail
	Name  string
}

type MenuItem struct {
	ID    MenuItemID
	Price float32
}

func New(customer Customer, menuItem MenuItem) (Order, error) {
	newOrder := Order{
		ID:            ID(uuid.NewString()),
		CustomerEmail: customer.Email,
		CustomerName:  customer.Name,
		MenuItemID:    menuItem.ID,
		Timestamps: []StatusUpdateTimestamp{
			{
				Status:    StatusOrdered,
				Timestamp: time.Now(),
			},
		},
		Status: StatusOrdered,
		Price:  menuItem.Price,
	}
	err := newOrder.Validate()
	if err != nil {
		return Order{}, fmt.Errorf("invalid order: %w", err)
	}
	return newOrder, nil
}
