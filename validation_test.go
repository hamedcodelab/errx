package errx

import (
	"errors"
	"log"
	"testing"
)

func TestValidationErr(t *testing.T) {
	log.Println(NewErrX().ValidationErr(errors.New("dfdf")))
}
