package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-git/go-git/v5"
)

var (
	App           = app.New()
	Window        = App.NewWindow("Packwiz GUI")
	Repository, _ = git.PlainOpen("C:\\Users\\oskar\\Documents\\darkmoon-modpack")
)

func main() {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			fmt.Println("Display help")
		}),
	)

	content := container.New(layout.NewCenterLayout(), widget.NewButtonWithIcon("Click me", theme.MediaPlayIcon(), func() {
		fmt.Println("Test")
	}))

	footer := container.NewCenter(widget.NewLabel("(c) 2023 Team Nyx"))

	Window.SetContent(container.NewBorder(toolbar, footer, nil, nil, content))
	Window.Resize(fyne.NewSize(1000, 600))
	Window.ShowAndRun()
}
