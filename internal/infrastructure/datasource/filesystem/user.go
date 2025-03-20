package filesystem

import (
	"fmt"
	"hackbar-copilot/internal/domain/user"
	"iter"
)

var _ user.Repository = (*userRepository)(nil)

type userRepository struct {
	fs    *filesystem
	index map[user.Email]int
}

func newUserRepository(fs *filesystem) *userRepository {
	index := make(map[user.Email]int)
	for i, user := range fs.data.users {
		index[user.Email] = i
	}
	return &userRepository{fs, index}
}

// All implements user.Repository.
func (u *userRepository) All() iter.Seq2[user.User, error] {
	return func(yield func(user.User, error) bool) {
		for _, user := range u.fs.data.users {
			if !yield(user, nil) {
				return
			}
		}
	}
}

// Get implements user.Repository.
func (u *userRepository) Get(email user.Email) (user.User, error) {
	i, ok := u.index[email]
	if !ok {
		return user.User{}, fmt.Errorf("user not found")
	}
	return u.fs.data.users[i], nil
}

// Save implements user.Repository.
func (u *userRepository) Save(d user.User) error {
	u.fs.data.users = append(u.fs.data.users, d)
	u.index[d.Email] = len(u.fs.data.users) - 1
	return nil
}
