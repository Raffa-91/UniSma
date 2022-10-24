package view

import (
	"UniSma/controller"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func tableWindow() fyne.CanvasObject {
	themeText := canvas.NewText("Theme", nil)
	settings := container.NewVBox(themeText)
	return settings
}
