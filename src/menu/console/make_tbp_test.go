package consolemenu

import (
	"formatore/src/enums"
	"formatore/src/io"
	"formatore/src/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeTableBlueprint(t *testing.T) {
	tableName := "tableName"
	q1 := "q1"
	t1 := enums.Text
	q2 := "q2"
	t2 := enums.Text
	input := []string{tableName, q1, t1, q2, t2, enums.DONE_TOKENS[0]}
	mockIO := &io.IO{
		I: &io.MockInput{Data: input},
		O: &io.MockOutput{},
	}

	cm := makeConsoleMenuWithIO(mockIO)

	cbs := []structs.ColumnBlueprint{{Name: q1, Type: t1}, {Name: q2, Type: t2}}
	expected := structs.TableBlueprint{
		Name:             tableName,
		ColumnBlueprints: cbs,
	}
	res, err := cm.MakeTableBlueprint()
	assert.NoError(t, err, "Should not error.")
	assert.Equal(t, expected, res.Content, "Should be equal.")
}
