package io

import (
	"formatore/enums"
	"fmt"
)

type OutputDisplay interface {
	Display(s string)
	ClearScreen()
	ClearThenDisplay(s string)
}

// Console output implementation
type FmtOutput struct{}
func (o *FmtOutput) Display(s string) {
	fmt.Print(s)
}

func (o *FmtOutput) ClearScreen() {
	fmt.Print(enums.ANSI_ESCAPE_CODE)
}

func (o *FmtOutput) ClearThenDisplay(s string) {
	o.ClearScreen()
	o.Display(s)
}

// Mock output implementation
type MockOutput struct {
	Data []string
}
func (m *MockOutput) Display(s string) {
	m.Data = append(m.Data, s)
}

func (m *MockOutput) ClearScreen() {
	m.Data = []string{}
}

func (m *MockOutput) ClearThenDisplay(s string) {
	m.ClearScreen()
	m.Display(s)
}
