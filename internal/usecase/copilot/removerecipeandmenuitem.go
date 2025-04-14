package copilot

// RemoveRecipeAndMenuItem implements Copilot.
func (c *copilot) RemoveRecipeAndMenuItem(name string) error {
	err := c.menu.Remove(name)
	if err != nil {
		return err
	}
	return c.recipe.Remove(name)
}
