package actions

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/nyxmc/packwiz-gui/ui"
)

func Add() {
	input := widget.NewEntry()

	ui.Content.RemoveAll()
	ui.Content.Add(input)

	input.OnChanged = func(v string) {
		fmt.Println(v)
	}
}
