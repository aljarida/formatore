package menu

import (
	"formatore/io"
)

type Menu interface {
	Render()
	Input() io.ResponseStatus
	Next() *Menu
}
