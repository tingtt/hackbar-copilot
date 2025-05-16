package filesystem

import (
	"hackbar-copilot/internal/usecase/barcounter"
	"hackbar-copilot/internal/usecase/cashier"
	"hackbar-copilot/internal/usecase/copilot"
	orderusecase "hackbar-copilot/internal/usecase/order"
	"sync"
)

func (f *filesystem) initializeGateway() {
	f.gateway = gateway{
		order:    orderRepository{f, &sync.RWMutex{}},
		checkout: checkoutRepository{f, &sync.RWMutex{}},
		cashout:  cashoutRepository{f, &sync.RWMutex{}},
		recipe:   recipeRepository{f, &sync.RWMutex{}},
		stock:    stockRepository{f, &sync.RWMutex{}},
		menu:     menuRepository{f, &sync.RWMutex{}},
		user:     userRepository{f, &sync.RWMutex{}},
	}
}

type gateway struct {
	order    orderRepository
	checkout checkoutRepository
	cashout  cashoutRepository
	recipe   recipeRepository
	stock    stockRepository
	menu     menuRepository
	user     userRepository
}

// BarCounterGateway implements Filesystem.
func (f *filesystem) BarCounterGateway() barcounter.Gateway {
	return &barcounterGateway{&f.gateway}
}

// CashierGateway implements Filesystem.
func (f *filesystem) CashierGateway() cashier.Gateway {
	return &cashierGateway{&f.gateway}
}

// CopilotGateway implements Filesystem.
func (f *filesystem) CopilotGateway() copilot.Gateway {
	return &copilotGateway{&f.gateway}
}

// OrderGateway implements Filesystem.
func (f *filesystem) OrderGateway() orderusecase.Gateway {
	return &orderGateway{&f.gateway}
}
