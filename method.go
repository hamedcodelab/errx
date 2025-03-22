package errx

import (
	"encoding/json"
	"net/http"
	"runtime"
)

type ErrX struct {
	Typ   Error
	Msg   string
	Code  int
	Stack string
	Err   error
}

func NewErrX(e error) ErrXModel {
	if e, ok := e.(ErrXModel); ok {
		return e
	}

	stackBuf := make([]byte, 1024)
	runtime.Stack(stackBuf, false)

	return &ErrX{
		Msg:   e.Error(),
		Err:   e,
		Stack: string(stackBuf),
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

func (e *ErrX) GetStructuredMessage() string {
	type StructuredMessage struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	msg := StructuredMessage{
		Message: e.Msg,
		Code:    e.Code,
	}

	b, _ := json.Marshal(msg)

	return string(b)
}

func (e *ErrX) GetCode() int {
	return e.Code
}

func (e *ErrX) GetStack() string {
	return e.Stack
}
