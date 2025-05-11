package consolemenu

import (
	"formatore/io"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMatchUserInput(t *testing.T) {
	cm := mockConsoleMenuWithStrArr([]string{})
	cm.options["testing"] = func() {}
	cm.initCharsToOptionNames()

	res, ok := cm.matchUserInput("t")
	assert.Equal(t, "testing", res)
	assert.Equal(t, true, ok)

	res, ok = cm.matchUserInput("testing")
	assert.Equal(t, "testing", res)
	assert.Equal(t, true, ok)

	res, ok = cm.matchUserInput("teasing")
	assert.Equal(t, "", res)
	assert.Equal(t, false, ok)
}

func TestInput(t *testing.T) {
	cm := mockConsoleMenuWithStrArr([]string{"t", "q", "d"})
	cm.options["testing"] = func() {}
	cm.initCharsToOptionNames()

	responseStatus := cm.Input()
	assert.Equal(t, io.InputOkay, responseStatus)
	assert.Equal(t, "t", cm.choice)


	cm.choice = ""
	responseStatus = cm.Input()
	assert.Equal(t, io.InputQuit, responseStatus)
	assert.Equal(t, "", cm.choice)

	cm.choice = ""
	responseStatus = cm.Input()
	assert.Equal(t, io.InputDone, responseStatus)
	assert.Equal(t, "", cm.choice)
}
