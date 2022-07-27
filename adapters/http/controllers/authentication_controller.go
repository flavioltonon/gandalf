package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/flavioltonon/gandalf/adapters/http/presenter"
	"github.com/flavioltonon/gandalf/application"
	"github.com/flavioltonon/gandalf/application/usecases"
	"github.com/flavioltonon/gandalf/common"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type AuthenticationController struct {
	authenticationService usecases.AuthenticationUsecase
	presenter             presenter.Presenter
	logger                common.Logger
}

func NewAuthenticationController(
	authenticationService usecases.AuthenticationUsecase,
	presenter presenter.Presenter,
	logger common.Logger,
) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: authenticationService,
		presenter:             presenter,
		logger:                logger,
	}
}

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *RegisterUserDTO) Validate() error {
	return ozzo.ValidateStruct(dto,
		ozzo.Field(&dto.Username, ozzo.Required),
		ozzo.Field(&dto.Password, ozzo.Required),
	)
}

func (c *AuthenticationController) RegisterUser(rw http.ResponseWriter, r *http.Request) {
	var dto RegisterUserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		return
	}

	if err := dto.Validate(); err != nil {
		c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		return
	}

	user, err := c.authenticationService.RegisterUser(r.Context(), dto.Username, dto.Password)
	if err != nil {
		switch {
		case application.IsValidationError(err), errors.Is(err, application.ErrUsernameAlreadyTaken):
			c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		default:
			c.presenter.Present(rw, http.StatusInternalServerError, presenter.NewError(err))
		}

		return
	}

	c.logger.Info(fmt.Sprintf("user created: %s", user.ID))

	c.presenter.Present(rw, http.StatusCreated, presenter.NewUser(user))
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginDTO) Validate() error {
	return ozzo.ValidateStruct(dto,
		ozzo.Field(&dto.Username, ozzo.Required),
		ozzo.Field(&dto.Password, ozzo.Required),
	)
}

func (c *AuthenticationController) Login(rw http.ResponseWriter, r *http.Request) {
	var dto LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		return
	}

	if err := dto.Validate(); err != nil {
		c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		return
	}

	user, err := c.authenticationService.Login(r.Context(), dto.Username, dto.Password)
	if err != nil {
		switch {
		case application.IsValidationError(err):
			c.presenter.Present(rw, http.StatusBadRequest, presenter.NewError(err))
		case errors.Is(err, application.ErrInvalidCredentials):
			c.presenter.Present(rw, http.StatusUnauthorized, presenter.NewError(err))
		default:
			c.presenter.Present(rw, http.StatusInternalServerError, presenter.NewError(err))
		}

		return
	}

	c.logger.Info(fmt.Sprintf("user logged in: %s", user.ID))

	c.presenter.Present(rw, http.StatusOK, presenter.NewUser(user))
}
