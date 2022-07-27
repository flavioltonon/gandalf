package usecases

import "github.com/flavioltonon/gandalf/domain/entity"

type AuthenticationUsecase interface {
	// RegisterUser registers a new user. In case the username has already been taken, an application.ErrUsernameAlreadyTaken
	// should be returned.
	RegisterUser(username, password string) (*entity.User, error)

	// Login authenticates a user. In case the credentials are invalid, an application.ErrInvalidCredentials should be returned.
	Login(username, password string) (*entity.User, error)
}
