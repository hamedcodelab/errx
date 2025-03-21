package errx

type ErrXModel interface {
	error
	ValidationErr(err error) ErrXModel
}

type ErrX struct {
	Typ Error
	Msg string
}

func NewErrX() ErrXModel {
	return &ErrX{}
}

func (e *ErrX) Error() string {
	switch e.Typ {
	case ErrorValidation:
		return ErrorValidation.String() + e.Msg
	default:
		return ""
	}
}

func (e *ErrX) ValidationErr(err error) ErrXModel {
	e.Typ = ErrorValidation
	e.Msg = err.Error()
	return e
}
