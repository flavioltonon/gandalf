package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/flavioltonon/gandalf/application"
	"github.com/flavioltonon/gandalf/domain"
	"github.com/flavioltonon/gandalf/domain/entity"
	"github.com/flavioltonon/gandalf/domain/interfaces"
	"github.com/flavioltonon/gandalf/domain/valueobject"
)

type AuthenticationService struct {
	usersRepository   interfaces.UsersRepository
	uuidFactory       interfaces.UUIDFactory
	passwordEncryptor interfaces.Encryptor
}

func NewAuthenticationService(
	usersRepository interfaces.UsersRepository,
	uuidFactory interfaces.UUIDFactory,
	passwordEncryptor interfaces.Encryptor) *AuthenticationService {
	return &AuthenticationService{usersRepository: usersRepository, uuidFactory: uuidFactory, passwordEncryptor: passwordEncryptor}
}

// RegisterUser registers a new user
func (s *AuthenticationService) RegisterUser(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := entity.NewUser(
		valueobject.NewID(s.uuidFactory.NewUUID()),
		valueobject.NewUsername(username),
		valueobject.NewPassword(s.passwordEncryptor.Encrypt(password)))
	if err != nil {
		return nil, fmt.Errorf("new user: %w", application.ValidationError(err))
	}

	if err := s.usersRepository.CreateUser(ctx, user); err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, application.ErrUsernameAlreadyTaken
		}

		return nil, fmt.Errorf("create user: %w", application.InternalError(err))
	}

	return user, nil
}

// Login authenticates a user
func (s *AuthenticationService) Login(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := s.usersRepository.GetUserByUsernameAndPassword(
		ctx,
		valueobject.NewUsername(username),
		valueobject.NewPassword(s.passwordEncryptor.Encrypt(password)))
	if errors.Is(err, domain.ErrNotFound) {
		return nil, application.ErrInvalidCredentials
	}
	if err != nil {
		return nil, fmt.Errorf("get user by username and password: %w", application.InternalError(err))
	}

	return user, nil
}
