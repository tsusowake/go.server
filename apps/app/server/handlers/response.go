package handlers

import "github.com/labstack/echo/v4"

type ResponseHandler struct{}

func (h *ResponseHandler) Respond(c echo.Context, req interface{}) error {
	return nil
}

func (h *ResponseHandler) RespondError(c echo.Context, req interface{}) error {
	return nil
}

func (h *ResponseHandler) respondJSON(c echo.Context, req interface{}) error {
	return nil
}
