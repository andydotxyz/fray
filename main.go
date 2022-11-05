package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fynedesk"
	systray "fyne.io/fynedesk/modules/systray"
)

func main() {
	a := app.NewWithID("xyz.andy.fray")
	w := a.NewWindow("Fray")
	w.SetPadded(false)

	fynedesk.SetInstance(&mockDesk{})

	m := systray.NewTray()
	w.SetContent(m.(fynedesk.StatusAreaModule).StatusAreaWidget())
	w.ShowAndRun()
}
