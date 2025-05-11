package consolemenu

import (
	"formatore/io"
	"log"
)

func (cm *ConsoleMenu) SetNext(next *ConsoleMenu) {
	cm.next = next
}

func (cm *ConsoleMenu) SetParent(parent *ConsoleMenu) {
	cm.parent = parent
}

func (cm *ConsoleMenu) SetHeaders(hs CMHeaders) {
	cm.headers = hs
}

func (cm *ConsoleMenu) SetHeaderTitle(s string) {
	cm.headers.Title = s
}

func (cm *ConsoleMenu) SetHeaderGuidance(s string) {
	cm.headers.Guidance = s
}

func (cm *ConsoleMenu) SetHeaderControls(s string) {
	cm.headers.Controls = s
}

func (cm *ConsoleMenu) SetHeaderError(s string) {
	cm.headers.Error = s
}

func (cm *ConsoleMenu) SetOptions(options map[string]func()) {
	log.Print("SetOptions found!")
	cm.options = options
	cm.initCharsToOptionNames()
}

func (cm *ConsoleMenu) AddOption(name string, fn func()) {
	cm.options[name] = fn
	cm.initCharsToOptionNames()
}

func (cm *ConsoleMenu) SetIO(io *io.IO) {
	cm.io = io
}
