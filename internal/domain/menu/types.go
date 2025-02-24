package menu

type Group struct {
	Name     string
	ImageURL *string
	Flavor   *string
	Items    []Item
}

type Item struct {
	Name       string
	ImageURL   *string
	Materials  []string
	OutOfStock bool
	Price      float32
}
