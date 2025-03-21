package errx

import (
	"errors"
	"log"
	"testing"
)

func TestValidationErr(t *testing.T) {
	log.Println(Forbidden(errors.New("test validation")).Error())
}
