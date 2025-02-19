package uilogic

import (
	"formatore/utils"
)

type IO struct {
	i InputReader
	o OutputDisplay
}

func (io *IO) getResponse(
		validator func(string) bool,
		prompt string,
		invalidMsg string) (string, error) {
	if validator == nil {
		return "", ErrNeedValidator
	}
	
	io.o.Display(prompt)
	for {
		response, err := io.i.Read()
		if err != nil {
			return "", err
		}
		
		valid := validator(response)
		if !valid {
			io.o.Display(invalidMsg)
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

func (io *IO) getQuestion() (utils.ColumnBlueprint, error) {
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

// TODO: Add unit tests.
func (io *IO) getQuestions() ([]utils.ColumnBlueprint, error) {
	questions := []utils.ColumnBlueprint{}
	for {
		question, err := io.getQuestion()
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
