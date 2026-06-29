package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/morikuni/failure/v2"

	pkgerror "github.com/tsusowake/go.server/pkg/error"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (c *CustomValidator) Validate(i any) error {
	if err := c.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

type BaseParams struct{}

func ValidateMiddleware[T any]() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			type Params struct {
				BaseParams
				Request T
			}
			var v *Params

			if err := c.Validate(v); err != nil {
				return failure.New(
					pkgerror.ErrorCodeBadRequest,
					failure.Message(err.Error()),
				)
			}
			if e := next(c); e != nil {
				return e
			}
			return nil
		}
	}
}
