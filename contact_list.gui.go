// auto-generated
// Code generated by GUI builder.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	cont        *fyne.Container
	contactList *widget.List
}

func newGUI() *gui {
	return &gui{}
}

func (g *gui) makeUI() fyne.CanvasObject {
	g.contactList = widget.NewList(func() int { return len(myList) }, func() fyne.CanvasObject {
		return container.New(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
	}, func(id widget.ListItemID, item fyne.CanvasObject) {
		item.(*fyne.Container).Objects[1].(*widget.Label).SetText(myList[id].Name + " " + myList[id].Email)
	})
	g.cont = container.NewMax(
		g.contactList)

	return g.cont
}
