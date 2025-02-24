package orderadapter

import (
	"hackbar-copilot/internal/domain/order"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
)

// StatusUpdateTimestamps implements OutputAdapter.
func (o *outputAdapter) StatusUpdateTimestamps(statusUpdates []order.StatusUpdateTimestamp) []*model.OrderStatusUpdateTimestamp {
	convertedStatusUpdates := []*model.OrderStatusUpdateTimestamp{}
	for _, statusUpdate := range statusUpdates {
		convertedStatusUpdates = append(convertedStatusUpdates, o.StatusUpdateTimestamp(statusUpdate))
	}
	return convertedStatusUpdates
}
