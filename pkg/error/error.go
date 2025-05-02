package error

const (
	// 4xx

	ErrorCodeBadRequest           = "BAD_REQUEST"
	ErrorCodeUnauthorized         = "UNAUTHORIZED"
	ErrorCodeForbidden            = "FORBIDDEN"
	ErrorCodeNotFound             = "NOT_FOUND"
	ErrorCodeConflict             = "CONFLICT"
	ErrorCodeUnprocessable        = "UNPROCESSABLE_ENTITY"
	ErrorCodePreconditionRequired = "PRECONDITION_REQUIRED"

	// 5xx

	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrorCodeBadGateway          = "BAD_GATEWAY"
	ErrorCodeServiceUnavailable  = "SERVICE_UNAVAILABLE"
	ErrorCodeGatewayTimeout      = "GATEWAY_TIMEOUT"
)
