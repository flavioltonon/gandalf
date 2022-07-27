package presenter

import "github.com/flavioltonon/gandalf/domain/entity"

type User struct {
	ID string `json:"_id"`
}

func NewUser(e *entity.User) *User {
	return &User{
		ID: e.ID.String(),
	}
}
