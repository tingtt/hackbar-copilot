package filesystem

import (
	"hackbar-copilot/internal/usecase/copilot"
)

var _ copilot.Gateway = (*copilotGateway)(nil)

type copilotGateway struct {
	*gateway
}

// Menu implements copilot.Gateway.
func (c *copilotGateway) Menu() copilot.MenuSaveListRemover {
	return &c.menu
}

// Recipe implements copilot.Gateway.
func (c *copilotGateway) Recipe() copilot.RecipeSaveListRemover {
	return &c.recipe
}

// Stock implements copilot.Gateway.
func (c *copilotGateway) Stock() copilot.StockSaveLister {
	return &c.stock
}
