package errors

import (
	"errors"
)

var ErrNeedValidator = errors.New("Validation function can not be nil.")
