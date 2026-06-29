package errutil

type ErrorCode int

const (
	ErrorCodeCanceled ErrorCode = iota + 1
	ErrorCodeUnknown
	ErrorCodeInvalidArgument
	ErrorCodeDeadlineExceeded
	ErrorCodeNotFound
	ErrorCodeAlreadyExists
	ErrorCodePermissionDenied
	ErrorCodeResourceExhausted
	ErrorCodeFailedPrecondition
	ErrorCodeAborted
	ErrorCodeOutOfRange
	ErrorCodeUnimplemented
	ErrorCodeInternal
	ErrorCodeUnavailable
	ErrorCodeDataLoss
	ErrorCodeUnauthenticated
)
