package ordersummary

import (
	"hackbar-copilot/internal/domain/order"
	"iter"
	"time"

	"github.com/tingtt/options"
)

type summarizer interface {
	Summarize(orders []order.Order) (Summary, error)
}

type ListerOption struct {
	Since *time.Time
}

func Since(t time.Time) options.Applier[ListerOption] {
	return func(lo *ListerOption) {
		lo.Since = &t
	}
}

type lister interface {
	Latest(optionAppliers ...options.Applier[ListerOption]) iter.Seq2[Summary, error]
}

type SummarizeLister interface {
	summarizer
	lister
}

type Repository interface {
	lister
	Save(s Summary) error
}

func NewSummarizeLister(or order.Repository, r Repository) SummarizeLister {
	return &summarizeLister{or, r}
}

type summarizeLister struct {
	orderRepository order.Repository
	Repository
}
