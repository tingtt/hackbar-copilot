package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

type InputAdapter interface {
	ApplyStatus(status model.OrderStatus) (order.Status, error)
}

func NewInputAdapter() InputAdapter {
	return &inputAdapter{}
}

type inputAdapter struct{}
