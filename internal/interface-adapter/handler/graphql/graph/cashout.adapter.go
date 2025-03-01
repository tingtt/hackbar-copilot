package graph

import (
	"hackbar-copilot/internal/domain/cashout"
	"hackbar-copilot/internal/interface-adapter/handler/graphql/graph/model"
	"time"
)

type cashout_ cashout.Cashout

func (c cashout_) apply() *model.Cashout {
	m := model.Cashout{
		Checkouts: checkouts_(c.Checkouts).apply(),
		Revenue:   float64(c.Revenue),
		Timestamp: c.Timestamp.UTC().Format(time.RFC3339),
		StaffID:   string(c.StaffID),
	}
	return &m
}

type cashouts_ []cashout.Cashout

func (c cashouts_) apply() []*model.Cashout {
	cashouts := make([]*model.Cashout, 0, len(c))
	for _, cashout := range c {
		cashouts = append(cashouts, cashout_(cashout).apply())
	}
	return cashouts
}
