package main

import (
	// "fmt"
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	highlightedPl := 0

	// Dummy data
	// TODO: replace with API calls later on
	playliststringtest := "wee duty mood,chill,poggers,kpop,jap weirdge songs,weeb,normie,rock,linkin park,nc25,bandori"
	playlistArr := strings.Split(playliststringtest, ",")

	// Declaration of UI Elements
	grid := tview.NewGrid()
	search := tview.NewTextView().SetText(" Search")
	playlists := tview.NewTextView()
	songs := tview.NewTextView()
	nowplaying := tview.NewTextView().SetText("Fireworks Festival\nRadwimps")

	// Setting up UI Elements
	grid.
		SetRows(1, 0, 3).
		SetColumns(30, 0, 20).
		SetBorders(true).
		// row index, col index, row span, col span, minheight, minwidth, bool focus
		// Layout for screens narrower than 100 cells (menu and side bar are hidden).
		AddItem(search, 0, 0, 1, 3, 0, 0, false).
		AddItem(nowplaying, 2, 0, 1, 3, 0, 0, false).
		AddItem(songs, 1, 1, 1, 2, 0, 0, false).
		AddItem(playlists, 0, 0, 0, 0, 0, 0, true).
		// Layout for screens wider than 100 cells.
		AddItem(playlists, 1, 0, 1, 1, 0, 100, true)

	grid.SetBackgroundColor(tcell.ColorNone)
	search.SetBackgroundColor(tcell.ColorNone)
	playlists.SetBackgroundColor(tcell.ColorNone)
	nowplaying.SetBackgroundColor(tcell.ColorNone)
	songs.SetBackgroundColor(tcell.ColorNone)


	// Set up playlist selector
	var sb strings.Builder
	for i, playlist := range playlistArr {
		sb.WriteString(fmt.Sprintf("[\"%d\"]%s[\"\"]\n", i, playlist))
	}

	// Key bindings for highlighting of playlists for selection
	playlists.SetRegions(true).SetText(sb.String())
	playlists.Highlight(strconv.Itoa(highlightedPl))
	playlists.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		// Up event
		case event.Rune() == 'j' || event.Key() == tcell.KeyDown:
			highlightedPl++
		// Down event
		case event.Rune() == 'k' || event.Key() == tcell.KeyUp:
			highlightedPl--
		case event.Rune() == ' ' || event.Key() == tcell.KeyEnter:
			// TODO selection event to show songs in main screen
		}

		// Go checks modulo by following the sign of a in a%b, python does the opposite
		if highlightedPl < 0 {
			highlightedPl = len(playlistArr) - 1
			playlists.Highlight(fmt.Sprintf("%d", highlightedPl))
		} else {
			playlists.Highlight(fmt.Sprintf("%d", highlightedPl%len(playlistArr)))
		}
		return event
	})


	if err := app.SetRoot(grid, true).SetFocus(playlists).Run(); err != nil {
		panic(err)
	}
}
