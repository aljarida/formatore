package ui

import (
	"errors"
)

// TODO: Move these into enums.
const ANSI_ESCAPE_CODE = "\033[H\033[J"

var QUIT_TOKENS = []string{"q", "quit"}
var DONE_TOKENS = []string{"d", "done"}
var BACK_TOKENS = []string{"b", "back"}

var ErrNeedValidator = errors.New("Validation function can not be nil.")
