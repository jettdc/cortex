package ui

import "github.com/rivo/tview"

var app *tview.Application

func InitUI() {
	app = tview.NewApplication()
}

func GetApp() *tview.Application {
	return app
}
