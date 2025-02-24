package ordersummary

import (
	"hackbar-copilot/internal/domain/order"
	"time"
)

// Summarize implements SummarizeLister.
func (s *summarizeLister) Summarize(orders []order.Order) (Summary, error) {
	summary := Summary{
		Orders:    orders,
		Revenue:   0,
		Timestamp: time.Now(),
	}
	for _, o := range orders {
		summary.Revenue += o.Price
	}

	err := s.Repository.Save(summary)
	if err != nil {
		return Summary{}, err
	}
	return summary, nil
}
