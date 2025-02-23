package main

import (
	"swim-pack-tool/tools"
	"swim-pack-tool/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	reg := tools.NewToolRegistry()
	tools.RegisterTools(reg)

	a := app.NewWithID("io.swimservices.swimpacktool")
	a.Settings().SetTheme(&forcedVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	w := a.NewWindow("Swim Pack Tool")
	w.SetMaster()

	content := container.NewStack()

	var currentTool tools.Tool
	w.SetOnDropped(func(p fyne.Position, u []fyne.URI) {
		if currentTool != nil {
			if len(u) > 0 {
				currentTool.OnDrop(u[0])
			}
		}
	})

	text := canvas.NewText("Swim Pack Tool", a.Settings().Theme().Color(theme.ColorNameForeground, a.Settings().ThemeVariant()))
	text.TextSize = 25

	split := container.NewBorder(nil, nil, container.New(utils.NewFixedSizeLayoutExpand(fyne.NewSize(200, 0)), container.NewBorder(text, nil, nil, nil, makeNav(func(t tools.Tool) {
		currentTool = t
		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}, reg))), nil, content)
	w.SetContent(split)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
