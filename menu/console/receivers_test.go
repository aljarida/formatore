package consolemenu

import (
	"formatore/enums"
	"formatore/structs"
	"formatore/io"
	"testing"
	"github.com/stretchr/testify/assert"
)

// TODO: Test MockOutput results.
func TestGetQuestion(t *testing.T) {
	expectedN := "What?" 
	expectedT := enums.Text
	mockIO := &io.IO{
		I: &io.MockInput{Data: []string{"NULL", "PRAGMA", "JOIN", expectedN, "SELECT", expectedT}},
		O: &io.MockOutput{},
	}
	cm := makeConsoleMenuWithIO(mockIO)

	res, err := cm.getQuestion()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, res.Content.Name, expectedN, "Should be equal.")
	assert.Equal(t, res.Content.Type, expectedT, "Should be equal.")

	cm.io.I = &io.MockInput{Data: []string{enums.QUIT_TOKENS[0], enums.DONE_TOKENS[0]}}

	res, err = cm.getQuestion()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
	
	res, err = cm.getQuestion()
	assert.NoError(t, err, "Should not error.")
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

	cm := makeConsoleMenuWithIO(mockIO)

	expected := []structs.ColumnBlueprint{
		{Name: q1, Type: t1},
		{Name: q2, Type: t2},
	}
	res, err := cm.GetQuestions()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")

	// Test set 2:
	setInputData(t, cm, []string{enums.DONE_TOKENS[0]})
	res, err = cm.GetQuestions()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputUserError, res.Status, "Should be equal.")
	
	// Test set 3:	
	input = []string{q1, t1, enums.QUIT_TOKENS[0]}
	setInputData(t, cm, input)
	res, err = cm.GetQuestions()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, io.InputQuit, res.Status, "Should be equal.")
}
