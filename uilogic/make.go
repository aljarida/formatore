package uilogic

import (
	"fmt"
	"formatore/dblogic"
	"formatore/structs"
	"strings"
)

func getResponse(prompt string,
				 invalid string,
				 validator func(string) bool,
			 	) (string, error) {
	if validator == nil {
		return "", fmt.Errorf("Validation function can not be nil.")
	}

	var response string
	for isValid := false; !isValid; {
		display(prompt)
		fmt.Scanln(&response)
		response = strings.ToUpper(response)
		isValid = validator(response)
		if !isValid {
			display(invalid)
		}
		if isCancel(response) {
			return "", ErrUserCanceled
		} 
	}

	return response, nil
}

func getQuestionType() (string, error) {
	return getResponse("Type:", "Invalid type.", dblogic.IsValidType)
}

var isNotReserved = func(s string) bool {
	return !dblogic.IsReserved(s)
}

func getQuestionText() (string, error) {
	return getResponse("Question:", "Invalid question.", isNotReserved)
}

func getQuestion() (structs.ColumnBlueprint, error) {
	qText, err := getQuestionText()
	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	qType, err := getQuestionType()
	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	return structs.ColumnBlueprint{Name: qText, Type: qType}, nil
}

func getQuestions() ([]structs.ColumnBlueprint, error) {
	more := true
	questions := make([]structs.ColumnBlueprint, 10)
	for more {
		quest, err := getQuestion()
		if err == ErrUserQuit {
			return []structs.ColumnBlueprint{}, err
		}
		if err == nil {
			questions = append(questions, quest)
		}
	}
	return questions, nil
}

func getTableName() (string, error) {
	return getResponse("Table name:", "Invalid table name.", isNotReserved)
}


// TODO: Analyse the possible error outputs and understand their flow.
// TODO: Consider removing "quit" and "cancel" errors, both or singularly.
// TODO: Add unit tests throughout this code.
// TODO: Abstract fmt.Scanln such that the code instead reads from an io.Reader.
func userMakeTable() (structs.TableBlueprint, error) {
	tableName, err := getTableName()
	if err != nil {
		return structs.TableBlueprint{}, err
	}
	
	questions, err := getQuestions()
	if err == ErrUserQuit { 
		return structs.TableBlueprint{}, err	
	} else if err != nil {
		return structs.TableBlueprint{}, fmt.Errorf("Error in userMakeTable() logic.")
	}

	tb := structs.TableBlueprint{
		Name: tableName,
		ColumnBlueprints: questions,
	}
	return tb, nil
}
