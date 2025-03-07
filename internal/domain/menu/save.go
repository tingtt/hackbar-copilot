package menu

import "fmt"

func (s *saveFindLister) Save(g Item) error {
	sanitizedGroup := g.Sanitized()
	if err := sanitizedGroup.Validate(); err != nil {
		return fmt.Errorf("invalid menu group: %w", err)
	}
	return s.Repository.Save(sanitizedGroup)
}
