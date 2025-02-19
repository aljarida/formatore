package uilogic

import (
	"testing"
	"formatore/utils"
	"github.com/stretchr/testify/assert"
)

func TestUiMakeTable(t *testing.T) {
	tableName := "tableName"
	q1 := "q1"
	t1 := utils.Text
	q2 := "q2"
	t2 := utils.Text
	input := []string{tableName, q1, t1, q2, t2, doneTokens[0]}
	io := &IO{
		i: &MockInput{input},
		o: &MockOutput{},
	}

	expected := utils.TableBlueprint{
		Name: tableName,
		ColumnBlueprints: []utils.ColumnBlueprint{{q1, t1}, {q2, t2}},
	}
	result, err := UiMakeTable(io)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, result, "Should be equal.")
}
