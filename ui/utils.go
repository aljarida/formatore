package ui

import (
	"formatore/utils"
	"strings"
	"unicode"
)

// TODO: Consider moving into global utils.

func isQuit(s string) bool {
	return utils.Has(QUIT_TOKENS, s)
}

func isDone(s string) bool {
	return utils.Has(DONE_TOKENS, s)
}

func prettyColumnNameAsQuestion(s string) string {
	var b strings.Builder
	b.Grow(len(s) + len("[?]"))

	first := true
	for _, r := range s {
		switch r {
		case '_':
			b.WriteRune(' ')
		default:
			if first {
				b.WriteRune(unicode.ToUpper(r))
				first = false
			} else {
				b.WriteRune(r)
			}
		}
	}

	b.WriteString("[?]")
	return b.String()
}

// Takes a string, returns parentheses around the first character.
func (cm *ConsoleMenu) parenthesizeFirstChar(s string) string {
	if len(s) < 1 {
		return s
	}

	extra := ""
	if len(s) > 1 {
		extra = s[1:]
	}

	return utils.JoinStrings("(", string(s[0]), ")", extra)
}

