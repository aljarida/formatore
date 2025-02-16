package uilogic

import (
	"errors"
)

const ansiEscapeCode = "\033[h\033[2J"
var cancelTokens = []string{"c"}
var quitTokens = []string{"q"}

var ErrUserCanceled = errors.New("User canceled.")
var ErrUserQuit = errors.New("User quit")
