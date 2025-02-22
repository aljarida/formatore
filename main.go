package main

import (
	"fmt"
	"formatore/uilogic"
	"formatore/utils"
)

func main() {
	io := &uilogic.IO{
		I: &uilogic.FmtInput{},
		O: &uilogic.FmtOutput{},
	}
	cb := []utils.ColumnBlueprint{
		{"What is your age?", utils.Integer},
		{"What is your name?", utils.Text},
		{"What do you want?", utils.Text},
	}
	res, err := uilogic.GetValues(io, cb)
	if err != nil {
		return
	}
	fmt.Print(res)
}
