package handlers

import (
	"github.com/tsusowake/go.server/domain"
)

type BaseHandler struct {
	Repository domain.Repository
	// Usecase
}

func NewBaseHandler(r domain.Repository) *BaseHandler {
	return &BaseHandler{
		Repository: r,
	}
}
