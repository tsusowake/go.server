package error

import (
	"net/http"

	"github.com/morikuni/failure/v2"
)

type APIErrorCode string

type APIError struct {
	StatusCode int32
	ErrorCode  APIErrorCode
	Message    string
	OrigErr    error
}

func NewAPIError(
	statusCode int32,
	message string,
	errorCode APIErrorCode,
	err error,
) error {
	return failure.New(APIError{
		StatusCode: statusCode,
		Message:    message,
		ErrorCode:  errorCode,
		OrigErr:    err,
	})
}

func (e APIError) Error() string {
	return e.OrigErr.Error()
}

func WithMessage(message string) func(*APIError) {
	return func(e *APIError) {
		e.Message = message
	}
}
func WithErrorCode(errorCode APIErrorCode) func(*APIError) {
	return func(e *APIError) {
		e.ErrorCode = errorCode
	}
}
func WithOrigErr(err error) func(*APIError) {
	return func(e *APIError) {
		e.OrigErr = err
	}
}

func NewBadRequestError(
	opts ...func(apiError *APIError),
) error {
	option := &APIError{
		Message:   "不正なリクエストです。",
		ErrorCode: "BAD_REQUEST",
		OrigErr:   failure.New("Bad Request"),
	}
	for _, opt := range opts {
		opt(option)
	}
	return NewAPIError(http.StatusBadRequest, option.Message, option.ErrorCode, option.OrigErr)
}

func NewConflictError(
	opts ...func(apiError *APIError),
) error {
	option := &APIError{
		Message:   "既に登録されています。",
		ErrorCode: "CONFLICT",
		OrigErr:   failure.New("Conflict"),
	}
	for _, opt := range opts {
		opt(option)
	}
	return NewAPIError(http.StatusConflict, option.Message, option.ErrorCode, option.OrigErr)
}

func NewServerError(
	opts ...func(apiError *APIError),
) error {
	option := &APIError{
		Message:   "サーバーエラーが発生しました。時間を空けてから再度お試しください。",
		ErrorCode: "INTERNAL_SERVER_ERROR",
		OrigErr:   failure.New("Internal Server Error"),
	}
	for _, opt := range opts {
		opt(option)
	}
	return NewAPIError(http.StatusInternalServerError, option.Message, option.ErrorCode, option.OrigErr)
}
