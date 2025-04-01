package recipe

type RecipeGroup struct {
	Name     string
	ImageURL *string
	Recipes  []Recipe
}

type Recipe struct {
	Name     string
	Category string
	Type     string // build, stir, shake etc.
	Glass    string // collins, shot, rock, beer etc.
	Steps    []Step
}

type Step struct {
	Material    *string
	Amount      *string
	Description *string
}

type RecipeType struct {
	Name        string
	Description *string
}

type GlassType struct {
	Name        string
	ImageURL    *string
	Description *string
}
