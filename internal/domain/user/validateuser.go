package user

import "fmt"

func (u User) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
