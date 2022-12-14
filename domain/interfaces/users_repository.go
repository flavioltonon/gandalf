package interfaces

import (
	"context"

	"github.com/flavioltonon/gandalf/domain/entity"
	"github.com/flavioltonon/gandalf/domain/valueobject"
)

type UsersRepository interface {
	// CreateUser persists a entity.User in the repository. In case it already exists, a domain.ErrAlreadyExists should be
	// returned instead.
	CreateUser(ctx context.Context, user *entity.User) error

	// GetUserByUsername returns an entity.User identified by a given username. In case it has not been created yet, a
	// domain.ErrNotFound should be returned instead.
	GetUserByUsernameAndPassword(ctx context.Context, username valueobject.Username, password valueobject.Password) (*entity.User, error)
}
