package response

import "github.com/tsusowake/go.server/internal/domain/auth/entity"

type User struct {
	ID string `json:"id"`
}

func ToUser(u *entity.User) *User {
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
