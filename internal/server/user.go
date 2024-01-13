package server

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/tsusowake/go.server/internal/server/request"
	"github.com/tsusowake/go.server/internal/server/response"
	"github.com/tsusowake/go.server/util/echoutil"
)

func (s *server) getUser(ctx echo.Context) error {
	req := new(request.GetUser)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	u, err := s.Database.User.GetByID(echoutil.FromEchoContext(ctx), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, "NotFound")
		}
		s.Logger.Error("getUser: error", zap.Error(err))
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(u))
}

func (s *server) createUser(ctx echo.Context) error {
	id, err := s.Database.User.Create(echoutil.FromEchoContext(ctx))
	if err != nil {
		s.Logger.Error("createUser: error", zap.Error(err))
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToCreateUserResponse(id))
}
