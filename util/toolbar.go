package util

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nyxmc/packwiz-gui/actions"
)

func UpdateToolbar(items ...widget.ToolbarItem) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ViewRefreshIcon(), actions.Refresh),
		widget.NewToolbarAction(theme.ContentAddIcon(), actions.Add),
	)

	if items != nil {
		for _, item := range items {
			toolbar.Append(item)
		}
	}

	toolbar.Append(widget.NewToolbarSpacer())
	toolbar.Append(widget.NewToolbarAction(theme.DownloadIcon(), actions.Pull))
	toolbar.Append(widget.NewToolbarAction(theme.UploadIcon(), actions.Upload))

	return toolbar
}
