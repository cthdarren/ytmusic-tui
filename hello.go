package main

import (
	// "fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Dummy data
	// TODO: replace with API calls later on
	playliststringtest := "wee duty mood,chill,poggers,kpop,jap weirdge songs,weeb,normie,rock,linkin park,nc25,bandori"
	playlistArr := strings.Split(playliststringtest, ",")

	// UI Elements
	search := tview.NewTextView().SetText(" Search")
	playlists := tview.NewList()

	nowplaying := tview.NewTextView().SetText("Fireworks Festival\nRadwimps")
	grid := tview.NewGrid().
		SetRows(1, 0, 3).
		SetColumns(30, 0, 20).
		SetBorders(true).
		AddItem(search, 0, 0, 1, 3, 0, 10, false).
		AddItem(nowplaying, 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(playlists, 0, 0, 0, 0, 0, 0, true)

	// Layout for screens wider than 100 cells.
	// row index, col index, row span, col span, minheight, minwidth
	grid.AddItem(playlists, 1, 0, 1, 1, 0, 100, true)

	for _, playlist := range playlistArr {
		playlists.AddItem(playlist, "", 0, func() { })
	}

	playlists.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey { // 	if event.Rune() == 'j' {
		if event.Rune() == 'j'{
			return tcell.NewEventKey(tcell.KeyDown, 0, 0)
		}
		if event.Rune() ==  'k'{
			return tcell.NewEventKey(tcell.KeyUp, 0, 0)
		}
		return event
	})

	if err := app.SetRoot(grid, true).SetFocus(playlists).Run(); err != nil {
		panic(err)
	}
}
