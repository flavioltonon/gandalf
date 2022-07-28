package entity

import (
	"github.com/flavioltonon/gandalf/domain/valueobject"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID       valueobject.ID
	Username valueobject.Username
	Password valueobject.Password
}

func NewUser(id valueobject.ID, username valueobject.Username, password valueobject.Password) (*User, error) {
	user := &User{
		ID:       id,
		Username: username,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (e *User) Validate() error {
	return ozzo.ValidateStruct(e,
		ozzo.Field(&e.ID, ozzo.Required),
		ozzo.Field(&e.Username, ozzo.Required),
		ozzo.Field(&e.Password, ozzo.Required),
	)
}
