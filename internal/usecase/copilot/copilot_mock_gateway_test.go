package copilot

var _ Gateway = new(MockGateway)

type MockGateway struct {
	menu   *MockMenu
	recipe *MockRecipe
	stock  *MockStock
}

// Menu implements Gateway.
func (m *MockGateway) Menu() MenuSaveListRemover {
	return m.menu
}

// Recipe implements Gateway.
func (m *MockGateway) Recipe() RecipeSaveListRemover {
	return m.recipe
}

// Stock implements Gateway.
func (m *MockGateway) Stock() StockSaveLister {
	return m.stock
}
