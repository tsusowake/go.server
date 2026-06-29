package response

import (
	"github.com/tsusowake/go.server/domain/auth/model"
)

type User struct {
	ID string `json:"id"`
}

func ToUser(u *model.User) *User {
	return &User{
		ID: u.ID,
	}
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

func ToCreateUserResponse(id string) *CreateUserResponse {
	return &CreateUserResponse{
		ID: id,
	}
}
