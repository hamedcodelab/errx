package errx

import "net/http"

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

func (e *ErrX) GetType() Error {
	return e.Typ
}

func (e *ErrX) WithCode(c int) ErrXModel {
	e.Code = c
	return e
}

func (e *ErrX) Error() string {
	return string(e.Typ) + " " + e.Msg + " (" + http.StatusText(e.Code) + ")"
}

func (e *ErrX) GetMessage() string {
	return e.Msg
}

func (e *ErrX) GetCode() int {
	return e.Code
}
