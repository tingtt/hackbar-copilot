package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResolver(t *testing.T) {
	t.Parallel()

	t.Run("will panic on not enough dependencies", func(t *testing.T) {
		t.Parallel()

		assert.Panics(t, func() {
			NewResolver(Dependencies{})
		})
	})

	t.Run("will insert local dependency", func(t *testing.T) {
		t.Parallel()

		// TODO: Refactor to make it testable
		// resolver := NewResolver(Dependencies{
		// 	Copilot:      copilot.New(copilot.Dependencies{}),
		// 	OrderService: order.New(order.Dependencies{}),
		// 	Cashier:      cashier.New(cashier.Dependencies{}),
		// }).(*Resolver)
		// assert.NotNil(t, resolver.recipeAdapter)
		// assert.NotNil(t, resolver.menuAdapter)
		// assert.NotNil(t, resolver.orderAdapter)
		// assert.NotNil(t, resolver.authAdapter)
	})
}
