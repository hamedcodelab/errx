package errx

import "net/http"

func Validation(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorValidation).WithCode(http.StatusBadRequest)
}

func Unauthorized(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorUnauthorized).WithCode(http.StatusUnauthorized)
}

func Forbidden(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorForbidden).WithCode(http.StatusForbidden)
}

func NotFound(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorNotFound).WithCode(http.StatusNotFound)
}

func Conflict(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorConflict).WithCode(http.StatusConflict)
}

func InternalServerError(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorInternal).WithCode(http.StatusInternalServerError)
}

func ServiceUnavailable(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorServiceUnavailable).WithCode(http.StatusServiceUnavailable)
}

func GatewayTimeout(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorGatewayTimeout).WithCode(http.StatusGatewayTimeout)
}
