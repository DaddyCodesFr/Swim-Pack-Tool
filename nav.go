package main

import (
	"swim-pack-tool/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makeNav(toolFunc func(tools.Tool), reg *tools.Registry) fyne.CanvasObject {
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return reg.Names()
		},
		IsBranch: func(uid string) bool {
			return uid == ""
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("a")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			if uid != "" {
				toolFunc(reg.ToolByName(uid))
			}
		},
	}
	tree.Select("Port")
	return tree
}
