package errx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error string

const (
	ErrorValidation         Error = "Validation Error:"
	ErrorUnauthorized       Error = "Unauthorized Error:"
	ErrorForbidden          Error = "Forbidden Error:"
	ErrorNotFound           Error = "Not Found Error:"
	ErrorConflict           Error = "Conflict Error:"
	ErrorInternal           Error = "Internal Server Error:"
	ErrorServiceUnavailable Error = "Service Unavailable Error:"
	ErrorGatewayTimeout     Error = "Gateway Timeout Error:"
	ErrorTooManyRequests    Error = "Too Many Requests Error:"
)

func (e Error) String() string {
	return string(e)
}

func ConvertGRPCToErrorType(err error) Error {
	switch status.Code(err) {
	case codes.NotFound:
		return ErrorNotFound
	case codes.InvalidArgument:
		return ErrorValidation
	case codes.PermissionDenied:
		return ErrorForbidden
	case codes.Unauthenticated:
		return ErrorUnauthorized
	case codes.ResourceExhausted:
		return ErrorTooManyRequests
	case codes.AlreadyExists:
		return ErrorConflict
	default:
		return ErrorInternal
	}
}
