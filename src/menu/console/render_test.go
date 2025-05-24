package consolemenu

import (
	"formatore/src/io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	cm := mockConsoleMenuWithStrArr([]string{})
	cm.Display("A")
	out := cm.io.O.(*io.MockOutput).Data
	assert.Equal(t, "A", out[0])
	cm.clearScreen()

	cm.Displayln("A")
	out = cm.io.O.(*io.MockOutput).Data
	assert.Equal(t, "A\n", out[0])
	cm.clearScreen()

	cm.Newline()
	out = cm.io.O.(*io.MockOutput).Data
	assert.Equal(t, "\n", out[0])
	cm.clearScreen()

	cm.options["E"] = func() {}
	cm.headers = CMHeaders{
		Title:    "A",
		Guidance: "B",
		Controls: "C",
		Error:    "D",
	}

	cm.Render()
	out = cm.io.O.(*io.MockOutput).Data
	assert.Equal(t, []string{"A\n", "B\n", "C\n", "Options:\n", "(E)\n"}, out)

	cm.errorState = true
	cm.Render()
	out = cm.io.O.(*io.MockOutput).Data
	assert.Equal(t, []string{"A\n", "B\n", "C\n", "D\n", "Options:\n", "(E)\n"}, out)
}
