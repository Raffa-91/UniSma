package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func overviewWindow() fyne.CanvasObject {
	//Sonos Wohzimmer Buttons
	/*
		playpause := widget.NewButtonWithIcon("Play", theme.MediaPlayIcon(), func() {

		})

		stop := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {

		})

		volumeup := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {

		})

		muteunmute := widget.NewButtonWithIcon("", theme.VolumeMuteIcon(), func() {

		})

		volumedown := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {

		})*/

	btn1 := widget.NewButton("Test1", func() {
		controller.TestLight()
	})

	//

	/*
		btn2 := widget.NewButton("Test2", func() {
			controller.TestLight()
		})
		btn3 := widget.NewButton("Test3", func() {
			controller.TestLight()
		})
		btn4 := widget.NewButton("Test4", func() {
			controller.TestLight()
		})*/

	lbl1 := widget.NewLabel("Überschrift")

	bottom := btn1
	right := lbl1

	content := container.New(layout.NewBorderLayout(nil, bottom, nil, right), lbl1, btn1)
	return container.NewGridWithColumns(3, container.NewVBox(content))
}
