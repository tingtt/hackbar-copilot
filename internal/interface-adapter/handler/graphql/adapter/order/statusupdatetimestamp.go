package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"time"
)

// StatusUpdateTimestamp implements OutputAdapter.
func (o *outputAdapter) StatusUpdateTimestamp(statusUpdate order.StatusUpdateTimestamp) *model.OrderStatusUpdateTimestamp {
	return &model.OrderStatusUpdateTimestamp{
		Status:    o.Status(statusUpdate.Status),
		Timestamp: statusUpdate.Timestamp.UTC().Format(time.RFC3339),
	}
}
