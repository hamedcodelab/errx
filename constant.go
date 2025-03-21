package errx

type Error string

const (
	ErrorValidation Error = "Validation Error:"
)

func (e Error) String() string {
	return string(e)
}
