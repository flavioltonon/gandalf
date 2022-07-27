package services

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/flavioltonon/gandalf/application"
	"github.com/flavioltonon/gandalf/domain"
	"github.com/flavioltonon/gandalf/domain/entity"
	"github.com/flavioltonon/gandalf/domain/repositories"
	"github.com/flavioltonon/gandalf/domain/valueobject"
)

type AuthenticationService struct {
	usersRepository repositories.UsersRepository
}

func NewAuthenticationService(usersRepository repositories.UsersRepository) *AuthenticationService {
	return &AuthenticationService{usersRepository: usersRepository}
}

// RegisterUser registers a new user
func (s *AuthenticationService) RegisterUser(username, password string) (*entity.User, error) {
	user := &entity.User{
		ID:       valueobject.NewID(),
		Username: valueobject.Username(username),
		Password: valueobject.Password(s.encrypt(password)),
	}

	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", application.ValidationError(err))
	}

	if err := s.usersRepository.CreateUser(user); err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, application.ErrUsernameAlreadyTaken
		}

		return nil, fmt.Errorf("create user: %w", application.InternalError(err))
	}

	return user, nil
}

// Login authenticates a user
func (s *AuthenticationService) Login(username, password string) (*entity.User, error) {
	user, err := s.usersRepository.GetUserByUsernameAndPassword(
		valueobject.Username(username),
		valueobject.Password(s.encrypt(password)),
	)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, application.ErrInvalidCredentials
	}
	if err != nil {
		return nil, fmt.Errorf("get user by username and password: %w", application.InternalError(err))
	}

	return user, nil
}

func (s *AuthenticationService) encrypt(value string) string {
	hash := md5.New()
	hash.Write([]byte(value))
	return string(hash.Sum(nil))
}
