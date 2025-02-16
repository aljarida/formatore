package uilogic

import (
	"fmt"
	"formatore/dblogic"
)

// TODO: Consider using Termbox to implement cards.
// NOTE: Consider how the architecture would change.

func getName() string {
	fmt.Println("Table name: ")
	var tableName string
	fmt.Scanln(&tableName)
	return tableName
}

func getQuestionType() (string, error) {
	fmt.Println("Type: ")
	var questionType string
	isValid := IsValidType(questionType)
	if !isValid {
		return "", fmt.Errorf("Invalid question type '%s'.", questionType)
	}
	fmt.Scanln(&questionType)
	return questionType, nil
}

func getQuestion() Question {
	fmt.Println("Question: ")
	var question string
	fmt.Scanln(question)
	getQuestionType()
}

func getQuestions() []Question {
	more := true	
	questions := make([]Question, 10)
	for more {
		q := getQuestion()
		questions = append(questions, q)
	}
	return questions
}

func userMakeTable() {
	tableName := getName()
	questions := getQuestions()
}
