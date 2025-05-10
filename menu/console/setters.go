package consolemenu

func (cm *ConsoleMenu) SetNext(next *ConsoleMenu) {
	cm.next = next
}

func (cm *ConsoleMenu) SetParent(parent *ConsoleMenu) {
	cm.parent = parent
}

func (cm *ConsoleMenu) SetHeaders(hs cmHeaders) {
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
