package errors

import (
	"errors"
)

var ErrNeedValidator = errors.New("Validation function can not be nil.")
var ErrTooManyInvalidResponses = errors.New("Too many invalid responses were logged, potentially indicating a testing error.")
