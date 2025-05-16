package filesystem

import (
	"hackbar-copilot/internal/domain/user"
	orderusecase "hackbar-copilot/internal/usecase/order"
	usecaseutils "hackbar-copilot/internal/usecase/utils"
	"iter"
	"sync"
)

var _ orderusecase.UserSaveGetter = (*userRepository)(nil)

type userRepository struct {
	fs    *filesystem
	mutex *sync.RWMutex
}

// All implements user.Repository.
func (r *userRepository) All() iter.Seq2[user.User, error] {
	return func(yield func(user.User, error) bool) {
		r.mutex.RLock()
		defer r.mutex.RUnlock()

		for _, user := range r.fs.data.users {
			if !yield(user, nil) {
				return
			}
		}
	}
}

// Get implements user.Repository.
func (r *userRepository) Get(email user.Email) (user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.fs.data.users {
		if user.Email == email {
			return user, nil
		}
	}
	return user.User{}, usecaseutils.ErrNotFound
}

// Save implements user.Repository.
func (r *userRepository) Save(d user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, user := range r.fs.data.users {
		if user.Email == d.Email {
			r.fs.data.users[i] = d
			return nil
		}
	}
	r.fs.data.users = append(r.fs.data.users, d)
	return nil
}
