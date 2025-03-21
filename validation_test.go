package errx

import (
	"errors"
	"log"
	"testing"
)

func TestValidationErr(t *testing.T) {
	log.Println(Validation(errors.New("test validation")).Error())
}
