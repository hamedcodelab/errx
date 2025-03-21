package errx

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
)

func (e Error) String() string {
	return string(e)
}
