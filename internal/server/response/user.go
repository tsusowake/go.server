package response

import "github.com/tsusowake/go.server/internal/domain/account/entity"

type User struct {
	ID uint64 `json:"id"`
}

func ToUser(u *entity.User) *User {
	return &User{
		ID: u.ID,
	}
}

type CreateUserResponse struct {
	ID uint64 `json:"id"`
}

func ToCreateUserResponse(id uint64) *CreateUserResponse {
	return &CreateUserResponse{
		ID: id,
	}
}
