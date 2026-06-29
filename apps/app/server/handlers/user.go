package handlers

import (
	"context"

	"github.com/tsusowake/go.server/apps/app/server/oapi"
)

// UserHandler implements the user-related operations of oapi.StrictServerInterface.
type UserHandler struct {
	*BaseHandler
}

func NewUserHandler(base *BaseHandler) *UserHandler {
	return &UserHandler{BaseHandler: base}
}

// CreateUser handles POST /users.
func (h *UserHandler) CreateUser(ctx context.Context, request oapi.CreateUserRequestObject) (oapi.CreateUserResponseObject, error) {
	// TODO: call usecase/repository with request.Body.Email
	return oapi.CreateUser201JSONResponse{
		Id:    "generated-id",
		Email: request.Body.Email,
	}, nil
}

// GetUser handles GET /users/{id}.
func (h *UserHandler) GetUser(ctx context.Context, request oapi.GetUserRequestObject) (oapi.GetUserResponseObject, error) {
	// TODO: fetch user by request.Id
	return oapi.GetUser200JSONResponse{
		Id:    request.Id,
		Email: "user@example.com",
	}, nil
}
