package ui

import (
	"errors"
)

// TODO: Move these into enums.
const ansiEscapeCode = "\033[h\033[2J"

var quitTokens = []string{"q"}
var doneTokens = []string{"d"}

var ErrUserQuit = errors.New("User quit.")
var ErrUserDone = errors.New("User done.")
var ErrNeedValidator = errors.New("Validation function can not be nil.")
