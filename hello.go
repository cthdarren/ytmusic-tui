package main

import (
	// "fmt"
	"example/hello/components"
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightUpAndLeft
	app := components.NewGui()

	// Dummy data
	// TODO: replace with API calls later on

	// Declaration of UI Elements
	search := tview.NewTextView()
	nowplaying := tview.NewTextView().SetText("Fireworks Festival\nRadwimps")

	// Make all elements transparent bg
	search.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Search").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	nowplaying.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Playing").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)


	// Setting up UI Elements


	if err := &app.app.SetRoot(grid, true).SetFocus(playlists).Run(); err != nil {
		panic(err)
	}
}
