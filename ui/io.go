package ui

import (
	"formatore/utils"
	"formatore/structs"
	"fmt"
)

type IO struct {
	I InputReader
	O OutputDisplay
}

func (io *IO) GetResponse(
		validator func(string) bool,
		prompt string,
		invalidMsg string) (string, error) {
	if validator == nil {
		return "", ErrNeedValidator
	}

	io.O.Display(prompt)
	for {
		response, err := io.I.Read()
		if err != nil {
			return "", err
		}
	
		valid := validator(response)
		if !valid {
			io.O.Display(invalidMsg)
			continue
		} else if isDone(response) {
			return "", ErrUserDone
		} else if isQuit(response) {
			return "", ErrUserQuit
		} else {
			return response, nil
		}
	}
}

// Move the functions into separate files (along with their unit tests).
func getQuestion(io *IO) (structs.ColumnBlueprint, error) {
	qText, err := io.GetResponse(utils.IsNotReserved, "Question:", "Invalid question.")
	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	qType, err := io.GetResponse(utils.IsValidType, "Type:", "Invalid type.")
	if err != nil {
		return structs.ColumnBlueprint{}, err
	}

	return structs.ColumnBlueprint{Name: qText, Type: qType}, nil
}

func getQuestions(io *IO) ([]structs.ColumnBlueprint, error) {
	questions := []structs.ColumnBlueprint{}
	for {
		question, err := getQuestion(io)
		if err == ErrUserDone && len(questions) > 0 {
 			return questions, nil
		} else if err == ErrUserDone { // && len(questions == 0)
			return []structs.ColumnBlueprint{}, ErrUserQuit
		} else if err != nil {
			return []structs.ColumnBlueprint{}, err
		} else {
			questions = append(questions, question)
		}
	}
}

// TODO: Add unit tests.
func GetValues(io *IO, cbs []structs.ColumnBlueprint) ([]string, error) {
	values := []string{}
	for i, cb := range cbs {
		prompt := fmt.Sprintf("%d. %s (%s):", i, cb.Name, cb.Type)
		invalidMsg := "Answer must match type."
		val := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		res, err := io.GetResponse(val, prompt, invalidMsg)
		if err != nil {
			return nil, err
		}

		values = append(values, res)
	}

	return values, nil
}
