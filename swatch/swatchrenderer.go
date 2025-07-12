package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)
type SwatchRenderer struct {
	square canvas.Rectangle
	objects []fyne.CanvasObject
	parent *Swatch
}
func (r *SwatchRenderer) Layout(size fyne.Size) {
	r.square.Resize(size)
	r.square.Move(fyne.NewPos(0, 0))
	for _, obj := range r.objects {
		obj.Resize(size)
		obj.Move(fyne.NewPos(0, 0))
	}
}