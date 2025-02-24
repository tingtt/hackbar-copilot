package orderadapter

type Adapter interface {
	OutputAdapter
	InputAdapter
}

func New() Adapter {
	return &orderAdapter{NewOutputAdapter(), NewInputAdapter()}
}

type orderAdapter struct {
	OutputAdapter
	InputAdapter
}
