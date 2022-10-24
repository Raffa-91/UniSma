package view

import (
	"UniSma/controller"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func kitchenWindow() fyne.CanvasObject {

	lblBrght := widget.NewLabel("Helligkeit")
	lblCol := widget.NewLabel("Color")
	splitter := widget.NewRichTextFromMarkdown(`---`)
	picSonos := canvas.NewImageFromResource(resources.GetResource(resources.SonosLogo, "SonosLogo"))
	picSonos.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		picSonos.SetMinSize(fyne.NewSize(112, 112))
	} else {
		picSonos.SetMinSize(fyne.NewSize(150, 150))
	}
	picLight := canvas.NewImageFromResource(resources.GetResource(resources.Light, "LightIcon"))
	picLight.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		picLight.SetMinSize(fyne.NewSize(112, 112))
	} else {
		picLight.SetMinSize(fyne.NewSize(150, 150))
	}

	//Licht Küche Haupt
	btnKitchenMain := widget.NewButton("On / Off", func() {

	})

	contKitchenMain := container.NewVBox(picLight, btnKitchenMain)

	//Licht Küche Schrank
	btnKitchenCab := widget.NewButton("On / Off", func() {

	})
	kitchenCabCol := 50.0
	sldKitchenCabColData := binding.BindFloat(&kitchenCabCol)
	sldKitchenCabCol := widget.NewSliderWithData(0, 100, sldKitchenCabColData)

	kitchenCabBrght := 50.0
	sldKitchenCabBrghtData := binding.BindFloat(&kitchenCabBrght)
	sldKitchenCabBrght := widget.NewSliderWithData(0, 100, sldKitchenCabBrghtData)

	contKitchenCab := container.NewVBox(picLight, btnKitchenCab, lblBrght, sldKitchenCabBrght, lblCol, sldKitchenCabCol)

	//Licht Küche Bar
	btnKitchenBar := widget.NewButton("On / Off", func() {

	})

	contKitchenBar := container.NewVBox(picLight, btnKitchenBar)

	//Musik Bar

	playpause := widget.NewButtonWithIcon("Play", theme.MediaPlayIcon(), func() {

	})

	stop := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {

	})

	volumeup := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {

	})

	muteunmute := widget.NewButtonWithIcon("", theme.VolumeMuteIcon(), func() {

	})

	volumedown := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {

	})

	selKitchenMusic := widget.NewSelect([]string{"SRF1", "Radio Vintage", "Radio Argovia"}, func(s string) {
		// s = Ausgabewert
	})

	actionbuttons := container.New(&ButtonsLayout{buttonHeight: 1.0}, playpause, volumedown, muteunmute, volumeup, stop)
	contKitchenMusic := container.NewVBox(picSonos, splitter, selKitchenMusic, actionbuttons)

	// Cards Kitchen View
	cardKitchenMain := widget.NewCard("Haupt", "", contKitchenMain)
	cardKitchenCab := widget.NewCard("Schrank", "", contKitchenCab)
	cardKitchenBar := widget.NewCard("Bar", "", contKitchenBar)
	cardKitchenMusic := widget.NewCard("Musik", "", contKitchenMusic)

	return container.NewGridWithColumns(2,
		container.NewVBox(splitter, cardKitchenMain, splitter, cardKitchenCab),
		container.NewVBox(splitter, cardKitchenBar, splitter, cardKitchenMusic),
	)
}
