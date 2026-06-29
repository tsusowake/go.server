package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/tsusowake/go.server/apps/app/server/oapi"
)

// API aggregates every operation handler so that it satisfies
// oapi.StrictServerInterface. Embed additional *XxxHandler here as the
// OpenAPI spec grows.
type API struct {
	*UserHandler
}

// compile-time guarantee that all spec operations are implemented.
var _ oapi.StrictServerInterface = (*API)(nil)

func NewAPI(base *BaseHandler) *API {
	return &API{
		UserHandler: NewUserHandler(base),
	}
}

// Register wires the generated routes onto the Echo instance.
func (a *API) Register(e *echo.Echo) {
	si := oapi.NewStrictHandler(a, nil)
	oapi.RegisterHandlers(e, si)
}
