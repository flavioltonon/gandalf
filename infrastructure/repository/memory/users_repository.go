package memory

import (
	"context"
	"sync"

	"github.com/flavioltonon/gandalf/domain"
	"github.com/flavioltonon/gandalf/domain/entity"
	"github.com/flavioltonon/gandalf/domain/valueobject"
)

type UsersRepository struct {
	users sync.Map
}

func NewUsersRepository() *UsersRepository {
	return new(UsersRepository)
}

func (r *UsersRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if _, err := r.GetUserByUsername(ctx, user.Username); err == nil {
		return domain.ErrAlreadyExists
	}

	r.users.Store(user.ID, user)
	return nil
}

func (r *UsersRepository) GetUserByUsername(ctx context.Context, username valueobject.Username) (*entity.User, error) {
	var user *entity.User

	r.users.Range(func(key, value interface{}) bool {
		u := value.(*entity.User)

		if u.Username != username {
			return true
		}

		user = u
		return false
	})

	if user == nil {
		return nil, domain.ErrNotFound
	}

	return user, nil
}

func (r *UsersRepository) GetUserByUsernameAndPassword(ctx context.Context, username valueobject.Username, password valueobject.Password) (*entity.User, error) {
	var user *entity.User

	r.users.Range(func(key, value interface{}) bool {
		u := value.(*entity.User)

		if u.Username != username {
			return true
		}

		if u.Password != password {
			return false
		}

		user = u
		return false
	})

	if user == nil {
		return nil, domain.ErrNotFound
	}

	return user, nil
}
