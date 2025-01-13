package menu

func (g Group) Sanitized() Group {
	sanitized := g
	sanitized.Items = make([]Item, 0, len(g.Items))

	if g.ImageURL != nil && *g.ImageURL == "" {
		sanitized.ImageURL = nil
	}
	if g.Flavor != nil && *g.Flavor == "" {
		sanitized.Flavor = nil
	}
	for _, item := range g.Items {
		sanitized.Items = append(sanitized.Items, item.Sanitized())
	}
	return sanitized
}
