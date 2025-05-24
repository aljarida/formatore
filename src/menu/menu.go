package menu

import (
	"formatore/src/io"
)

type Menu interface {
	Render()
	Input() io.ResponseStatus
	Next() *Menu
}
