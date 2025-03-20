package user

import "fmt"

func (s *saveLister) Save(u User) error {
	if err := u.Validate(); err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}
	return s.Repository.Save(u)
}
