package errx

type ErrXModel interface {
	error
	WithType(typ Error) ErrXModel
	WithCode(c int) ErrXModel
	GetType() Error
	GetMessage() string
	GetStructuredMessage() string
	GetCode() int
	GetStack() string
}
