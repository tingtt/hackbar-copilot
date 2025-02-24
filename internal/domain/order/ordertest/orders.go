package ordertest

import (
	"hackbar-copilot/internal/domain/order"
	"time"
)

var ExampleOrders = []order.Order{
	{
		ID:         "1",
		CustomerID: "user@example.test",
		MenuItemID: order.MenuItemID{
			GroupName: "Phuket Sling",
			ItemName:  "Cocktail",
		},
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Date(2025, 2, 24, 21, 30, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusPrepared,
				Timestamp: time.Date(2025, 2, 24, 21, 32, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusDelivered,
				Timestamp: time.Date(2025, 2, 24, 21, 32, 30, 00, time.UTC),
			},
			{
				Status:    order.StatusCheckedOut,
				Timestamp: time.Date(2025, 2, 24, 23, 00, 00, 00, time.UTC),
			},
		},
		Status: order.StatusCheckedOut,
		Price:  700,
	},
	{
		ID:         "2",
		CustomerID: "user@example.test",
		MenuItemID: order.MenuItemID{
			GroupName: "Passoamoni",
			ItemName:  "Cocktail",
		},
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Date(2025, 2, 24, 22, 00, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusPrepared,
				Timestamp: time.Date(2025, 2, 24, 22, 02, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusDelivered,
				Timestamp: time.Date(2025, 2, 24, 22, 2, 30, 00, time.UTC),
			},
			{
				Status:    order.StatusCheckedOut,
				Timestamp: time.Date(2025, 2, 24, 23, 00, 00, 00, time.UTC),
			},
		},
		Status: order.StatusCheckedOut,
		Price:  700,
	},
	{
		ID:         "3",
		CustomerID: "user@example.test",
		MenuItemID: order.MenuItemID{
			GroupName: "Phuket Sling",
			ItemName:  "Cocktail",
		},
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Date(2025, 2, 25, 21, 30, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusPrepared,
				Timestamp: time.Date(2025, 2, 25, 21, 32, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusDelivered,
				Timestamp: time.Date(2025, 2, 25, 21, 32, 30, 00, time.UTC),
			},
		},
		Status: order.StatusDelivered,
		Price:  700,
	},
	{
		ID:         "4",
		CustomerID: "user@example.test",
		MenuItemID: order.MenuItemID{
			GroupName: "Passoamoni",
			ItemName:  "Cocktail",
		},
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Date(2025, 2, 25, 22, 00, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusPrepared,
				Timestamp: time.Date(2025, 2, 25, 22, 02, 00, 00, time.UTC),
			},
			{
				Status:    order.StatusDelivered,
				Timestamp: time.Date(2025, 2, 25, 22, 2, 30, 00, time.UTC),
			},
		},
		Status: order.StatusDelivered,
		Price:  700,
	},
	{
		ID:         "5",
		CustomerID: "user@example.test",
		MenuItemID: order.MenuItemID{
			GroupName: "Blue Devil",
			ItemName:  "Cocktail",
		},
		Timestamps: []order.StatusUpdateTimestamp{
			{
				Status:    order.StatusOrdered,
				Timestamp: time.Date(2025, 2, 25, 22, 30, 00, 00, time.UTC),
			},
		},
		Status: order.StatusOrdered,
		Price:  700,
	},
}

var ExampleOrdersIter = IterWithNilError(ExampleOrders)
