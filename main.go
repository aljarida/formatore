package main

import "formatore/src/app"

func main() {
	application := app.NewApp()
	application.Loop()
}
