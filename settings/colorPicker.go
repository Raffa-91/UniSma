package settings

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strings"
	"sync"
)

type themedBackground struct {
	widget.BaseWidget
}

type dialogLayout struct {
	d *dialog
}

type dialog struct {
	callback    func(bool)
	title       string
	icon        fyne.Resource
	desiredSize fyne.Size

	win            *widget.PopUp
	bg             *themedBackground
	content, label fyne.CanvasObject
	dismiss        *widget.Button
	parent         fyne.Window
	layout         *dialogLayout
}

type colorAdvancedPicker struct {
	widget.BaseWidget
	Red, Green, Blue, Alpha int // Range 0-255
	Hue                     int // Range 0-360 (degrees)
	Saturation, Lightness   int // Range 0-100 (percent)
	ColorModel              string

	onChange func(color.Color)
}

type ColorPickerDialog struct {
	*dialog
	Advanced bool
	color    color.Color
	callback func(c color.Color)
	advanced *widget.Accordion
	picker   *colorAdvancedPicker
}

type colorButton struct {
	widget.BaseWidget
	color   color.Color
	onTap   func(color.Color)
	hovered bool
}

func (c colorButton) CreateRenderer() fyne.WidgetRenderer {
	//TODO implement me
	panic("implement me")
}

type BaseWidget struct {
	size     fyne.Size
	position fyne.Position
	Hidden   bool

	impl         fyne.Widget
	propertyLock sync.RWMutex
}

func (w *BaseWidget) super() fyne.Widget {
	w.propertyLock.RLock()
	impl := w.impl
	w.propertyLock.RUnlock()
	return impl
}

func (w *BaseWidget) ExtendBaseWidget(wid fyne.Widget) {
	impl := w.super()
	if impl != nil {
		return
	}

	w.propertyLock.Lock()
	defer w.propertyLock.Unlock()
	w.impl = wid
}

// newColorButton creates a colorButton with the given color and callback.
func newColorButton(color color.Color, onTap func(color.Color)) *colorButton {
	b := &colorButton{
		color: color,
		onTap: onTap,
	}
	b.ExtendBaseWidget(b)
	return b
}

func stringToColor(s string) (color.Color, error) {
	var c color.NRGBA
	var err error
	if len(s) == 7 {
		c.A = 0xFF
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	} else {
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	}
	return c, err
}

func stringsToColors(ss ...string) (colors []color.Color) {
	for _, s := range ss {
		if s == "" {
			continue
		}
		c, err := stringToColor(s)
		if err != nil {
			fyne.LogError("Couldn't parse color:", err)
		} else {
			colors = append(colors, c)
		}
	}
	return
}

func newColorButtonBox(colors []color.Color, icon fyne.Resource, callback func(color.Color)) fyne.CanvasObject {
	var objects []fyne.CanvasObject
	if icon != nil && len(colors) > 0 {
		objects = append(objects, widget.NewIcon(icon))
	}
	for _, c := range colors {
		objects = append(objects, newColorButton(c, callback))
	}
	return container.NewGridWithColumns(8, objects...)
}

const (
	preferenceRecents    = "color_recents"
	preferenceMaxRecents = 8
)

func readRecentColors() (recents []string) {
	for _, r := range strings.Split(fyne.CurrentApp().Preferences().String(preferenceRecents), ",") {
		if r != "" {
			recents = append(recents, r)
		}
	}
	return
}

// newColorRecentPicker returns a component for selecting recent colors.
func newColorRecentPicker(callback func(color.Color)) fyne.CanvasObject {
	return newColorButtonBox(stringsToColors(readRecentColors()...), theme.HistoryIcon(), callback)
}
