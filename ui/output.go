package ui

import (
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
	fmt.Print(ansiEscapeCode)
}

func (o *FmtOutput) ClearThenDisplay(s string) {
	o.ClearScreen()
	o.Display(s)
}

// Mock output implementation
type MockOutput struct {}
func (m *MockOutput) Display(_ string) {}

func (m *MockOutput) ClearScreen() {}

func (m *MockOutput) ClearThenDisplay(_ string) {}
