package consolemenu

import (
	"formatore/src/enums"
	"formatore/src/errors"
	"formatore/src/io"
	"formatore/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoopUntilValidResponse(t *testing.T) {
	expected := "VALID"
	mockIO := io.IO{
		I: &io.MockInput{Data: []string{expected}},
		O: &io.MockOutput{},
	}
	cm := makeConsoleMenuWithIO(&mockIO)

	emptyHs := CMHeaders{}
	res, err := cm.loopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")
	assert.Equal(t, io.InputOkay, res.Status, "Should be equal.")

	setInputData(t, cm, []string{"NULL", expected})
	res, err = cm.loopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")

	_, err = cm.loopUntilValidResponse(nil, emptyHs)
	assert.Equal(t, errors.ErrNeedValidator, err, "Should be equal.")

	setInputData(t, cm, []string{enums.QUIT_TOKENS[0]})
	res, err = cm.loopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")

	setInputData(t, cm, []string{enums.DONE_TOKENS[0]})
	res, err = cm.loopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputDone, res.Status, "Should be equal.")

	setInputData(t, cm, make([]string, errorThreshold+1))
	notEmpty := func(s string) bool { return s != "" }
	_, err = cm.loopUntilValidResponse(notEmpty, emptyHs)
	assert.ErrorIs(t, errors.ErrTooManyInvalidResponses, err)

	setInputData(t, cm, []string{"a"})
	empty := func(s string) bool { return s == "" }
	testHeaders := CMHeaders{
		Guidance: "X",
		Error:    "Y",
	}
	_, _ = cm.loopUntilValidResponse(empty, testHeaders)
	assert.Equal(t, "X", cm.headers.Guidance)
	assert.Equal(t, "Y", cm.headers.Error)
}
