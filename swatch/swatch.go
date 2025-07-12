package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/apptype"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(swatch *Swatch)
}

func (swatch *Swatch) SetColor(c *color.Color) {
	swatch.Color = *c
	swatch.Refresh()
}
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	swatch := &Swatch{
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
		Selected:     false,
	}
	swatch.ExtendBaseWidget(swatch)
	return swatch
	
}
