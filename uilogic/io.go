package uilogic

import (
	"formatore/utils"
	"fmt"
)

type IO struct {
	I InputReader
	O OutputDisplay
}

func (io *IO) getResponse(
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
func getQuestion(io *IO) (utils.ColumnBlueprint, error) {
	qText, err := io.getResponse(utils.IsNotReserved, "Question:", "Invalid question.")
	if err != nil {
		return utils.ColumnBlueprint{}, err
	}

	qType, err := io.getResponse(utils.IsValidType, "Type:", "Invalid type.")
	if err != nil {
		return utils.ColumnBlueprint{}, err
	}

	return utils.ColumnBlueprint{Name: qText, Type: qType}, nil
}

func getQuestions(io *IO) ([]utils.ColumnBlueprint, error) {
	questions := []utils.ColumnBlueprint{}
	for {
		question, err := getQuestion(io)
		if err == ErrUserDone && len(questions) > 0 {
 			return questions, nil
		} else if err == ErrUserDone { // && len(questions == 0)
			return []utils.ColumnBlueprint{}, ErrUserQuit
		} else if err != nil {
			return []utils.ColumnBlueprint{}, err
		} else {
			questions = append(questions, question)
		}
	}
}

// TODO: Add unit tests.
func GetValues(io *IO, cbs []utils.ColumnBlueprint) ([]string, error) {
	values := []string{}
	for i, cb := range cbs {
		prompt := fmt.Sprintf("%d. %s (%s):", i, cb.Name, cb.Type)
		invalidMsg := "Answer must match type."
		val := func(s string) bool {
			return utils.InferType(s) == cb.Type
		}

		res, err := io.getResponse(val, prompt, invalidMsg)
		if err != nil {
			return nil, err
		}

		values = append(values, res)
	}

	return values, nil
}
