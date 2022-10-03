package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"os"
)

type ButtonsLayout struct {
	buttonHeight float32
}

func Start() {
	a := app.NewWithID("UniSma -Home Control")
	w := a.NewWindow("UniSma - Home Control")

	tabs := container.NewAppTabs(
		container.NewTabItem("Übericht", container.NewPadded(overviewWindow())),
		container.NewTabItem("Wohnzimmer", container.NewPadded(livingroomWindow())),
		container.NewTabItem("Küche", container.NewPadded(kitchenWindow())),
		container.NewTabItem("Tisch", container.NewPadded(tableWindow())),
		container.NewTabItem("Settings", container.NewPadded(settingsWindow())),
		container.NewTabItem("About", aboutWindow()),
	)

	tabs.OnSelected = func(t *container.TabItem) {
		t.Content.Refresh()
	}

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(w.Canvas().Size().Width, w.Canvas().Size().Height*0.6))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
	os.Exit(0)
}

/*
func Gui() {
	a := app.New()

	w := a.NewWindow("UniSma")

	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(), tutorial)

	split.Offset = 0
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}*/
