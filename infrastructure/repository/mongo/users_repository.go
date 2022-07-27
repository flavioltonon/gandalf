package mongo

import (
	"context"
	"errors"

	"github.com/flavioltonon/gandalf/domain"
	"github.com/flavioltonon/gandalf/domain/entity"
	"github.com/flavioltonon/gandalf/domain/valueobject"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func newUser(e *entity.User) User {
	return User{
		ID:       e.ID.String(),
		Username: e.Username.String(),
		Password: e.Password.String(),
	}
}

func (u *User) toEntity() *entity.User {
	return &entity.User{
		ID:       valueobject.ID(u.ID),
		Username: valueobject.Username(u.Username),
		Password: valueobject.Password(u.Password),
	}
}

type UsersRepository struct {
	users *Collection
}

func (d *Database) NewUsersRepository() *UsersRepository {
	return &UsersRepository{users: d.Collection("users")}
}

func (r *UsersRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.users.NewTransaction(ctx, func(sc mongo.SessionContext) (interface{}, error) {
		if err := r.users.FindOne(sc, bson.M{"username": user.Username}, new(User)); err != mongo.ErrNoDocuments {
			if err != nil {
				return nil, err
			}

			return nil, domain.ErrAlreadyExists
		}

		return r.users.collection.InsertOne(ctx, newUser(user))
	})
}

func (r *UsersRepository) GetUserByUsernameAndPassword(ctx context.Context, username valueobject.Username, password valueobject.Password) (*entity.User, error) {
	var user User

	if err := r.users.FindOne(ctx, bson.M{"username": username.String(), "password": password.String()}, &user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}

		return nil, err
	}

	return user.toEntity(), nil
}
