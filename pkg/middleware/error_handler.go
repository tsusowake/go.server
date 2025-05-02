package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morikuni/failure/v2"

	"github.com/tsusowake/go.server/pkg/context"
	pkgerror "github.com/tsusowake/go.server/pkg/error"
)

type ErrorResponse struct {
	ErrorCode pkgerror.APIErrorCode `json:"error_code"`
	Message   string                `json:"message"`
}

type Response struct {
	Data  interface{}    `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

func ErrorHandler(e error, c echo.Context) {
	if e != nil && !c.Response().Committed {
		var apiError pkgerror.APIError
		var err error

		switch {
		case errors.As(e, &apiError):
			err = respondWithJSON(c,
				int(apiError.StatusCode),
				&Response{
					Error: &ErrorResponse{
						ErrorCode: apiError.ErrorCode,
						Message:   apiError.Message,
					},
				},
			)
		default:
			// NOTE: その他のエラーはサーバーエラーとして扱う
			logError(c, e)
			err = respondWithJSON(c,
				http.StatusInternalServerError,
				&Response{
					Error: &ErrorResponse{
						ErrorCode: pkgerror.ErrorCodeInternalServerError,
						Message:   "サーバーエラーが発生しました。",
					},
				},
			)
		}

		if err != nil {
			logError(c, err)
		}
	}
}

func logError(c echo.Context, e error) {
	slog.ErrorContext(
		c.Request().Context(),
		"an error occurred",
		slog.String("error_code", fmt.Sprintf("%v", failure.CodeOf(e))),
		slog.String("message", e.Error()),
		slog.String("stack_trace", fmt.Sprintf("%v", failure.CallStackOf(e))),
		slog.String("cause", fmt.Sprintf("%v", failure.CauseOf(e))),
		slog.String("trace_id", context.GetTraceIDFrom(c.Request().Context())),
	)
}

func respondWithJSON(c echo.Context, statusCode int, response *Response) error {
	return failure.Wrap(c.JSON(statusCode, response))
}
