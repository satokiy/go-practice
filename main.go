package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)


func total(n int) int {
	if n == 0 {
		return 0
	}
	return n + total(n-1)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		widget.NewVBox(
			l,e,
			widget.NewButton("Hi!", func() {
				n, _ := strconv.Atoi(e.Text)
				l.SetText("Total " + strconv.Itoa(total(n)) + "!")
			}),
		),
	)
	a.Settings().SetTheme(theme.DarkTheme())
	w.ShowAndRun()
}
