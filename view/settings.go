package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func settingsWindow() fyne.CanvasObject {
	themeText := canvas.NewText("Theme", nil)
	dropdown := widget.NewSelect([]string{"Light", "Dark", "Default"}, parseTheme())
	theme := fyne.CurrentApp().Preferences().StringWithFallback("Theme", "Default")
	switch theme {
	case "Light":
		dropdown.PlaceHolder = "Light"
	case "Dark":
		dropdown.PlaceHolder = "Dark"
	case "Default":
		dropdown.PlaceHolder = "Default"
	}

	dropdown.Refresh()

	settings := container.NewVBox(themeText, dropdown)
	return settings

}

func parseTheme() func(string) {
	return func(t string) {
		switch t {
		case "Light":
			fyne.CurrentApp().Preferences().SetString("Theme", "Light")
			fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		case "Dark":
			fyne.CurrentApp().Preferences().SetString("Theme", "Dark")
			fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		case "Default":
			fyne.CurrentApp().Preferences().SetString("Theme", "Default")
			fyne.CurrentApp().Settings().SetTheme(theme.DefaultTheme())
		}
	}
}
