package errors

import (
	"errors"
)

var ErrNeedValidator = errors.New("Validation function can not be nil.")
var ErrTooManyInvalidResponses = errors.New("Too many invalid responses were logged, potentially indicating a testing error.")
var ErrTooFewColumns = errors.New("There were fewer than two columns found when finding a table's ColumnBlueprints, implying that the ID or time was missing.")
