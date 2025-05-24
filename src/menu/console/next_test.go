package consolemenu

import (
	"formatore/src/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNext(t *testing.T) {
	test_int := 0
	test_int_ptr := &test_int

	cm := mockConsoleMenuWithStrArr([]string{"t", "n", "d", "q"})
	cm.options["testing"] = func() { *test_int_ptr = 1 }
	cm.initCharsToOptionNames()

	cm.Input()
	next := cm.Next()
	assert.Equal(t, 1, test_int)
	assert.Equal(t, cm, next)

	new_cm := mockConsoleMenuWithStrArr([]string{})
	cm.options["new"] = func() { cm.SetNext(new_cm) }
	cm.initCharsToOptionNames()

	cm.Input()
	next = cm.Next()
	assert.Equal(t, new_cm, next)
	next = new_cm.Next()
	assert.Equal(t, new_cm, next)

	res := cm.Input()
	assert.Equal(t, cm.choice, "")
	assert.Equal(t, io.InputDone, res)

	res = cm.Input()
	assert.Equal(t, io.InputQuit, res)
}
