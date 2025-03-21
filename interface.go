package errx

import "net/http"

type ErrXModel interface {
	error
	WithType(typ Error) ErrXModel
	WithCode(c int) ErrXModel
}

type ErrX struct {
	Typ  Error
	Msg  string
	Code int
}

func NewErrX(e error) ErrXModel {
	if e, ok := e.(ErrXModel); ok {
		return e
	}

	return &ErrX{
		Msg: e.Error(),
	}
}

func (e *ErrX) WithType(typ Error) ErrXModel {
	e.Typ = typ
	return e
}

func (e *ErrX) WithCode(c int) ErrXModel {
	e.Code = c
	return e
}

func (e *ErrX) Error() string {
	switch e.Typ {
	case ErrorValidation:
		return ErrorValidation.String() + e.Msg + http.StatusText(e.Code)
	default:
		return ""
	}
}

func Validation(er error) ErrXModel {
	return NewErrX(er).WithType(ErrorValidation).WithCode(http.StatusBadRequest)
}
