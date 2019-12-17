package errors

import (
	"errors"
	"fmt"
)

func CreateError(message string, err error) error {
	return errors.New(fmt.Sprintf("%s %s", message, err.Error()))
}
