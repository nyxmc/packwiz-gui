package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nyxmc/packwiz-gui/util"
)

var (
	App     = app.New()
	Window  = App.NewWindow("Packwiz GUI")
	Content = container.NewCenter()
	Toolbar = util.UpdateToolbar()
	Footer  = container.NewCenter(widget.NewLabel("(c) 2023 Team Nyx"))
)

func Render() {
	Window.SetContent(container.NewBorder(Toolbar, Footer, nil, nil, Content))
}
