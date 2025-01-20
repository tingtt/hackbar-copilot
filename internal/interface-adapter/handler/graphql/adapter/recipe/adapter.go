package recipeadapter

type Adapter interface {
	InputAdapter
	OutputAdapter
}

func New() Adapter {
	return &recipeAdapter{NewInputAdapter(), NewOutputAdapter()}
}

type recipeAdapter struct {
	InputAdapter
	OutputAdapter
}
