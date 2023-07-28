package main

import (
	"fyne.io/fyne/v2"
	"github.com/nyxmc/packwiz-gui/ui"
)

func main() {
	ui.Render()

	ui.Window.Resize(fyne.NewSize(1000, 600))
	ui.Window.ShowAndRun()
}
