package ui

import (
	"formatore/utils"
)

func isQuit(s string) bool {
	return utils.Has(quitTokens, s)
}

func isDone(s string) bool {
	return utils.Has(doneTokens, s)
}
