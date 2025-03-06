package ui

import (
	"testing"
	"formatore/enums"
	"formatore/structs"
	"github.com/stretchr/testify/assert"
)

func TestMakeTable(t *testing.T) {
	tableName := "tableName"
	q1 := "q1"
	t1 := enums.Text
	q2 := "q2"
	t2 := enums.Text
	input := []string{tableName, q1, t1, q2, t2, doneTokens[0]}
	io := &IO{
		I: &MockInput{input},
		O: &MockOutput{},
	}

	expected := structs.TableBlueprint{
		Name: tableName,
		ColumnBlueprints: []structs.ColumnBlueprint{{q1, t1}, {q2, t2}},
	}
	result, err := MakeTable(io)
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, result, "Should be equal.")
}
