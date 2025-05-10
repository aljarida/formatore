package consolemenu

import (
	"formatore/enums"
	"formatore/utils"
	"formatore/structs"
	"formatore/errors"
	"formatore/io"
	"testing"
	"github.com/stretchr/testify/assert"
)

// TODO: Test error results.
// TODO: Test MockOutput results.
// TODO: Place this in a better location.
func makeMockConsoleMenu(ioObj *io.IO) *ConsoleMenu {
	return InitConsoleMenu(
		&ConsoleMenu{
			io: ioObj,
		},
	)
}

func TestLoopUntilValidResponse(t *testing.T) {
	expected := "VALID"
	mockIO := io.IO{
		I: &io.MockInput{Data: []string{expected}},
		O: &io.MockOutput{},
	}
	cm := makeMockConsoleMenu(&mockIO)

	emptyHs := cmHeaders{}
	res, _ := cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.Equal(t, expected, res.Content, "Should be equal.")

	cm.io.I.SetData([]string{"NULL", expected})
	res, _ = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.Equal(t, expected, res.Content, "Should be equal.")

	_, err := cm.LoopUntilValidResponse(nil, emptyHs)
	assert.Equal(t, errors.ErrNeedValidator, err, "Should be equal.")

	cm.io.I = &io.MockInput{Data: []string{enums.QUIT_TOKENS[0]}}
	res, _ = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
	
	cm.io.I = &io.MockInput{Data: []string{enums.DONE_TOKENS[0]}}
	res, _ = cm.LoopUntilValidResponse(utils.IsNotReserved, emptyHs)
	assert.Equal(t, io.InputDone, res.Status, "Should be equal.")
}

func TestGetQuestion(t *testing.T) {
	expectedN := "What?" 
	expectedT := enums.Text
	mockIO := &io.IO{
		I: &io.MockInput{Data: []string{"NULL", "PRAGMA", "JOIN", expectedN, "SELECT", expectedT}},
		O: &io.MockOutput{},
	}
	cm := makeMockConsoleMenu(mockIO)

	res, err := cm.getQuestion()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, res.Content.Name, expectedN, "Should be equal.")
	assert.Equal(t, res.Content.Type, expectedT, "Should be equal.")

	cm.io.I = &io.MockInput{Data: []string{enums.QUIT_TOKENS[0], enums.DONE_TOKENS[0]}}

	res, _ = cm.getQuestion()
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
	
	res, _ = cm.getQuestion()
	assert.Equal(t, io.InputDone, res.Status, "Should be equal.")
}

func TestGetQuestions(t *testing.T) {
	// Test set 1:
	q1 := "q1"
	t1 := enums.Text
	q2 := "q2"
	t2 := enums.Text
	input := []string{q1, t1,
					  q2, t2,
					  enums.DONE_TOKENS[0]}
	mockIO := &io.IO{
		I: &io.MockInput{Data: input},
		O: &io.MockOutput{},
	}

	cm := makeMockConsoleMenu(mockIO)

	expected := []structs.ColumnBlueprint{
		{Name: q1, Type: t1},
		{Name: q2, Type: t2},
	}
	res, err := cm.GetQuestions()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")

	// Test set 2:
	mockIO.I = &io.MockInput{Data: []string{enums.DONE_TOKENS[0]}}
	res, _ = cm.GetQuestions()
	assert.Equal(t, io.InputDone, res.Status, "Should be equal.")
	
	// Test set 3:	
	input = []string{q1, t2,
					 enums.QUIT_TOKENS[0]}
	cm.io.I = &io.MockInput{Data: input}
	res, _ = cm.GetQuestions()
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
}
