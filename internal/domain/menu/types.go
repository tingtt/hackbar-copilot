package menu

type Item struct {
	Name     string
	ImageURL *string
	Flavor   *string
	Options  []ItemOption
}

type ItemOption struct {
	Name       string
	ImageURL   *string
	Materials  []string
	OutOfStock bool
	Price      float32
}
