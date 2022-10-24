package view

import (
	"UniSma/controller"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func musicButtons() fyne.CanvasObject {
	return musicButtons()
}

func livingroomWindow() fyne.CanvasObject {

	//Config Pictures / Const Labels
	//
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

	// Light bubble living Room
	//
	btnLivingBub := widget.NewButton("On / Off", func() {

	})

	btnLivingBubColor := widget.NewButton("Color", func() {

		d := dialog.NewColorPicker("Farbe", "Wähle eine Farbe", func(c color.Color) {

		}, w)
		d.Show()

	})
	//livingBubCol := 50.0
	//sldLivingBubColData := binding.BindFloat(&livingBubCol)
	//sldLivingBubCol := widget.NewSliderWithData(0, 100, sldLivingBubColData)

	livingBubBrght := 50.0
	sldLivingBubBrghtData := binding.BindFloat(&livingBubBrght)
	sldLivingBubBrght := widget.NewSliderWithData(0, 100, sldLivingBubBrghtData)
	//Container Living Light Strip
	contLivingBub := container.NewVBox(picLight, btnLivingBub, lblBrght, sldLivingBubBrght, lblCol, btnLivingBubColor)

	// Light Strip Living Room
	//
	btnLivingLightStrip := widget.NewButton("On / Off", func() {

	})
	livingLightStripCol := 50.0
	sldLivingLightStripColData := binding.BindFloat(&livingLightStripCol)
	sldLivingLightStripCol := widget.NewSliderWithData(0, 100, sldLivingLightStripColData)

	livingLightStripBrght := 50.0
	sldLivingLightStripBrghtData := binding.BindFloat(&livingLightStripBrght)
	sldLivingLightStripBrght := widget.NewSliderWithData(0, 100, sldLivingLightStripBrghtData)
	//Container Living Light Strip
	contLivingLight := container.NewVBox(picLight, btnLivingLightStrip, lblBrght, sldLivingLightStripBrght, lblCol, sldLivingLightStripCol)

	//Music Living Room
	//
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
	//Container Living Music
	contLivingMusic := container.NewVBox(picSonos, splitter, selKitchenMusic, actionbuttons)

	//Cards Living Room
	//
	cardLightLiving := widget.NewCard("Licht", "Hue", contLivingLight)
	cardSonos := widget.NewCard("Musik", "Sonos", contLivingMusic)
	cardLivingBub := widget.NewCard("Kugel", "Hue", contLivingBub)

	return container.NewGridWithColumns(2, container.NewVBox(cardSonos, cardLightLiving), container.NewVBox(cardLivingBub))

}
