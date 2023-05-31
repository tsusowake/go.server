package request

import (
	"github.com/labstack/echo/v4"
)

type GetUser struct {
	ID string `json:"id" param:"id"`
}

type CreateUser struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func NewCreateUser(ctx echo.Context) (*CreateUser, error) {
	// FIXME validator の struct tag が効いてない
	req := new(CreateUser)
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}
	return req, nil
}
