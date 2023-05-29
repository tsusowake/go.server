package server

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/tsusowake/go.server/internal/server/request"
	"github.com/tsusowake/go.server/internal/server/response"
	"go.uber.org/zap"
	"net/http"
)

func (s *server) getUser(ctx echo.Context) error {
	req := new(request.GetUser)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	u, err := s.Database.User.GetByID(FromEchoContext(ctx), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, "NotFound")
		}
		s.Logger.Error("getUser: error", zap.Error(err))
		return err
	}
	return ctx.JSON(http.StatusOK, response.ToUser(u))
}
