package apptype

import( "fyne.io/fyne/v2")
// AppType defines the type of application being run.
type BrushType = int
type PxCanvasConfig struct{
	// Name is the name of the application.
	DrawingArea fyne.Size
	CanvasOffset fyne.Position
	PxRows , PxCols int
	// BrushType defines the type of brush used in the application.
	PxSize int
	
}