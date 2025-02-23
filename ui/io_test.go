package ui

import (
	"formatore/enums"
	"formatore/utils"
	"formatore/structs"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {
	expected := "VALID"
	io := IO{
		I: &MockInput{[]string{expected}},
		O: &MockOutput{},
	}

	result, _ := io.GetResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, expected, result, "Should be equal.")

	io.I = &MockInput{[]string{"NULL", expected}}
	result, _ = io.GetResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, expected, result, "Should be equal.")

	_, err := io.GetResponse(nil, "", "")
	assert.Equal(t, ErrNeedValidator, err, "Should be equal.")

	io.I = &MockInput{[]string{quitTokens[0]}}
	_, err = io.GetResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
	
	io.I = &MockInput{[]string{doneTokens[0]}}
	_, err = io.GetResponse(utils.IsNotReserved, "", "")
	assert.Equal(t, ErrUserDone, err, "Should be equal.")
}

func TestGetQuestion(t *testing.T) {
	expectedN := "What?" 
	expectedT := enums.Text
	io := &IO{
		I: &MockInput{[]string{"NULL", "PRAGMA", "JOIN", expectedN, "SELECT", expectedT}},
		O: &MockOutput{},
	}

	result, err := getQuestion(io)	
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, result.Name, expectedN, "Should be equal.")
	assert.Equal(t, result.Type, expectedT, "Should be equal.")

	io = &IO{
		I: &MockInput{[]string{quitTokens[0], doneTokens[0]}},
		O: &MockOutput{},
	}

	result, err = getQuestion(io)
	assert.Equal(t, err, ErrUserQuit, "Should be equal.")
	
	result, err = getQuestion(io)
	assert.Equal(t, err, ErrUserDone, "Should be equal.")
}

func TestGetQuestions(t *testing.T) {
	// Test set 1:
	q1 := "q1"
	t1 := enums.Text
	q2 := "q2"
	t2 := enums.Text
	input := []string{q1, t1,
					  q2, t2,
					  doneTokens[0]}
	io := &IO{
		I: &MockInput{input},
		O: &MockOutput{},
	}

	expected := []structs.ColumnBlueprint{
		{Name: q1, Type: t1},
		{Name: q2, Type: t2},
	}
	result, err := getQuestions(io)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, result, "Should be equal.")

	// Test set 2:
	input = []string{doneTokens[0]}
	io = &IO{
		I: &MockInput{input},
		O: &MockOutput{},
	}
	_, err = getQuestions(io)
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
	
	// Test set 3:	
	input = []string{q1, t2,
					 quitTokens[0]}
	io = &IO{
		I: &MockInput{input},
		O: &MockOutput{},
	}
	_, err = getQuestions(io)
	assert.Equal(t, ErrUserQuit, err, "Should be equal.")
}
