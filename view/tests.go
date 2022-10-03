package view

import (
	"fyne.io/fyne/v2"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Tests() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
		text1, text2, text3)
	myWindow.SetContent(grid)

	myWindow.Resize(fyne.NewSize(380, 75))
	myWindow.ShowAndRun()

}
