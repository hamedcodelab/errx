package errx

import (
	"errors"
	"log"
	"testing"
)

func TestValidationErr(t *testing.T) {
	log.Println(NewErrX(errors.New("dfdf")).ValidationErr())
}
