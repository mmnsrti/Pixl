package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

// AppType defines the type of application being run.
type BrushType = int
type PxCanvasConfig struct{
	// Name is the name of the application.
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows, PxCols int
	PxSize int
	
} 
type State struct{
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
	
}
func (s *State) SetFilePath(path string) {
	s.FilePath = path
}
