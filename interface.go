package errx

import "net/http"

type ErrXModel interface {
	error
	ValidationErr() ErrXModel
}

type ErrX struct {
	Typ  Error
	Msg  string
	Code int
}

func NewErrX(e any) ErrXModel {
	if e, ok := e.(ErrXModel); ok {
		return e
	}

	if e, ok := e.(error); ok {
		return newErrX(0, e.Error())
	}

	return &ErrX{}
}

func newErrX(code int, err any) ErrXModel {
	return &ErrX{Code: code, Msg: err.(string)}
}

func (e *ErrX) Error() string {
	switch e.Typ {
	case ErrorValidation:
		return ErrorValidation.String() + e.Msg + http.StatusText(e.Code)
	default:
		return ""
	}
}

func (e *ErrX) ValidationErr() ErrXModel {
	e.Typ = ErrorValidation
	e.Code = http.StatusUnprocessableEntity
	return e
}
