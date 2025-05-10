package consolemenu

import (
	"formatore/enums"
	"formatore/utils"
	"formatore/errors"
	"formatore/io"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSanityCheck(t *testing.T) {
    assert.True(t, true, "Sanity check should pass")
}

func TestLoopUntilValidResponse(t *testing.T) {
	expected := "VALID"
	mockIO := io.IO{
		I: &io.MockInput{Data: []string{expected}},
		O: &io.MockOutput{},
	}
	cm := makeMockConsoleMenu(&mockIO)

	emptyHs := cmHeaders{}
	res, err := cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")

	setInputData(t, cm, []string{"NULL", expected}) 
	res, err = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")

	_, err = cm.LoopUntilValidResponse(nil, emptyHs)
	assert.Equal(t, errors.ErrNeedValidator, err, "Should be equal.")

	setInputData(t, cm, []string{enums.QUIT_TOKENS[0]})
	res, err = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
	
	setInputData(t, cm, []string{enums.DONE_TOKENS[0]})
	res, err = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputDone, res.Status, "Should be equal.")
}
