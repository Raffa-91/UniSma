package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func tableWindow() fyne.CanvasObject {
	themeText := canvas.NewText("Theme", nil)
	settings := container.NewVBox(themeText)
	return settings
}
