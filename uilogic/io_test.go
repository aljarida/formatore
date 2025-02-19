package uilogic

import (
	"formatore/utils"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {
	expected := "VALID"
	io := IO{
		i: &MockInput{[]string{expected}},
		o: &MockOutput{},
	}

	result, _ := io.getResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, expected, result, "Should be equal.")

	io.i = &MockInput{[]string{"NULL", expected}}
	result, _ = io.getResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, expected, result, "Should be equal.")

	_, err := io.getResponse(nil, "", "")
	assert.Equal(t, ErrNeedValidator, err, "Should be equal.")

	io.i = &MockInput{[]string{quitTokens[0]}}
	_, err = io.getResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
	
	io.i = &MockInput{[]string{doneTokens[0]}}
	_, err = io.getResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, ErrUserDone, err, "Should be equal.")
}

func TestGetQuestion(t *testing.T) {
	expectedN := "What?" 
	expectedT := utils.Text
	io := IO{
		i: &MockInput{[]string{"NULL", "PRAGMA", "JOIN", expectedN, "SELECT", expectedT}},
		o: &MockOutput{},
	}

	result, err := io.getQuestion()	
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, result.Name, expectedN, "Should be equal.")
	assert.Equal(t, result.Type, expectedT, "Should be equal.")

	io = IO{
		i: &MockInput{[]string{quitTokens[0], doneTokens[0]}},
		o: &MockOutput{},
	}

	result, err = io.getQuestion()
	assert.Equal(t, err, ErrUserQuit, "Should be equal.")
	
	result, err = io.getQuestion()
	assert.Equal(t, err, ErrUserDone, "Should be equal.")
}

func TestGetQuestions(t *testing.T) {
	// Test set 1:
	q1 := "q1"
	t1 := utils.Text
	q2 := "q2"
	t2 := utils.Text
	input := []string{q1, t1,
					  q2, t2,
					  doneTokens[0]}
	io := IO{
		i: &MockInput{input},
		o: &MockOutput{},
	}

	expected := []utils.ColumnBlueprint{
		{q1, t1},
		{q2, t2},
	}
	result, err := io.getQuestions()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, result, "Should be equal.")

	// Test set 2:
	input = []string{doneTokens[0]}
	io = IO{
		i: &MockInput{input},
		o: &MockOutput{},
	}
	_, err = io.getQuestions()
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
	
	// Test set 3:	
	input = []string{q1, t2,
					 quitTokens[0]}
	io = IO{
		i: &MockInput{input},
		o: &MockOutput{},
	}
	_, err = io.getQuestions()
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
}
