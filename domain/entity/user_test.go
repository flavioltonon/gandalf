package entity

import (
	"testing"

	"github.com/flavioltonon/gandalf/domain/valueobject"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	var (
		id       = valueobject.NewID("84a92cd8-e1dd-4987-bd32-cbdde7588ff9")
		username = valueobject.NewUsername("username")
		password = valueobject.NewPassword("password")
	)

	type args struct {
		id       valueobject.ID
		username valueobject.Username
		password valueobject.Password
	}

	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Given valid parameters, a new User should be created",
			args: args{
				id:       id,
				username: username,
				password: password,
			},
			want: &User{
				ID:       id,
				Username: username,
				Password: password,
			},
		},
		{
			name: "If an invalid ID is provided, an error should be returned",
			args: args{
				username: username,
				password: password,
			},
			wantErr: true,
		},
		{
			name: "If an invalid username is provided, an error should be returned",
			args: args{
				id:       id,
				password: password,
			},
			wantErr: true,
		},
		{
			name: "If an invalid password is provided, an error should be returned",
			args: args{
				id:       id,
				username: username,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.username, tt.args.password)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
