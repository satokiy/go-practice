package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type cdata struct {
	memory int
	calc   string
	flag   bool
}

// Number buttons
func createNumberButtons(f func(v int)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		widget.NewButton(strconv.Itoa(1), func() { f(1) }),
		widget.NewButton(strconv.Itoa(2), func() { f(2) }),
		widget.NewButton(strconv.Itoa(3), func() { f(3) }),
		widget.NewButton(strconv.Itoa(4), func() { f(4) }),
		widget.NewButton(strconv.Itoa(5), func() { f(5) }),
		widget.NewButton(strconv.Itoa(6), func() { f(6) }),
		widget.NewButton(strconv.Itoa(7), func() { f(7) }),
		widget.NewButton(strconv.Itoa(8), func() { f(8) }),
		widget.NewButton(strconv.Itoa(9), func() { f(9) }),
		widget.NewButton(strconv.Itoa(0), func() { f(0) }),
	)
	return c
}

// Calc buttons
func createCalcButtons(f func(v string)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewButton(("CL"), func() { f("CL") }),
		widget.NewButton(("+"), func() { f("+") }),
		widget.NewButton(("-"), func() { f("-") }),
		widget.NewButton(("*"), func() { f("*") }),
		widget.NewButton(("/"), func() { f("/") }),
	)
	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("This is Calculator")
	l := widget.NewLabel("0")
	l.Alignment = fyne.TextAlignTrailing

	data := cdata{
		memory: 0,
		calc:   "",
		flag:   false,
	}

	calc := func(n int) {
		switch data.calc {
		case "+":
			data.memory += n
		case "-":
			data.memory -= n
		case "*":
			data.memory *= n
		case "/":
			data.memory /= n
		default:
			data.memory = n
		}
		l.SetText(strconv.Itoa(data.memory))
		data.flag = true
	}

	pushNum := func(v int) {
		s := l.Text
		if data.flag {
			s = "0"
			data.flag = false
		}

		s += strconv.Itoa(v)
		n, err := strconv.Atoi(s)
		if err == nil {
			l.SetText(strconv.Itoa(n))
		}
	}

	pushCalc := func(v string) {
		if v == "CL" {
			l.SetText("0")
			data.memory = 0
			data.calc = ""
			data.flag = false
			return
		}
		n, er := strconv.Atoi(l.Text)
		if er != nil {
			return
		}
		calc(n)
		data.calc = v
	}

	pushEnter := func() {
		n, er := strconv.Atoi(l.Text)
		if er != nil {
			return
		}
		calc(n)
		data.calc = ""
	}

	k := createNumberButtons(pushNum)
	c := createCalcButtons(pushCalc)
	e := widget.NewButton("Enter", pushEnter)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
			l, e, nil, c,
		),
			l,e,k,c,
		),
	)
	w.Resize(fyne.NewSize(400, 400))
	a.Settings().SetTheme(theme.DarkTheme())
	w.ShowAndRun()

}
