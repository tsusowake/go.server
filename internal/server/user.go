package server

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/tsusowake/go.server/internal/server/request"
	"github.com/tsusowake/go.server/internal/server/response"
	"github.com/tsusowake/go.server/pkg/echoutil"
)

func (s *server) getUser(ctx echo.Context) error {
	req := new(request.GetUser)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	u, err := s.Database.Auth.User.GetByID(echoutil.FromEchoContext(ctx), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, "NotFound")
		}
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(u))
}

func (s *server) createUser(ctx echo.Context) error {
	id, err := s.Database.Auth.User.Create(echoutil.FromEchoContext(ctx))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToCreateUserResponse(id))
}
