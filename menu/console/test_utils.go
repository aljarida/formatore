package consolemenu

import (
	"formatore/io"
	"testing"
)

func mockConsoleMenuWithStrArr(data []string) *ConsoleMenu {
	return InitConsoleMenu(
		&ConsoleMenu{
			io: &io.IO{
				I: &io.MockInput{Data: data},
				O: &io.MockOutput{},
			},
		},
	)
}

func makeConsoleMenuWithIO(ioObj *io.IO) *ConsoleMenu {
	return InitConsoleMenu(
		&ConsoleMenu{
			io: ioObj,
		},
	)
}

func setInputData(t *testing.T, cm *ConsoleMenu, data []string) {
	if mock, ok := cm.io.I.(*io.MockInput); ok {
		mock.SetData(data)
	} else {
		t.Fatalf("Expected cm.io.I to be a *MockInput; got %T.", cm.io.I)
	}
}
