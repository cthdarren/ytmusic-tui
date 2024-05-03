package main

import (
	// "fmt"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

func main() {
	tview.Borders.HorizontalFocus = tview.BoxDrawingsLightHorizontal
	tview.Borders.VerticalFocus = tview.BoxDrawingsLightVertical
	tview.Borders.TopLeftFocus = tview.BoxDrawingsLightDownAndRight
	tview.Borders.TopRightFocus = tview.BoxDrawingsLightDownAndLeft
	tview.Borders.BottomLeftFocus = tview.BoxDrawingsLightUpAndRight
	tview.Borders.BottomRightFocus = tview.BoxDrawingsLightUpAndLeft
	app := tview.NewApplication()
	highlightedPl := 0

	// Dummy data
	// TODO: replace with API calls later on
	playliststringtest := "wee duty mood,chill,poggers,kpop,jap weirdge songs,weeb,normie,rock,linkin park,nc25,bandori"
	playlistArr := strings.Split(playliststringtest, ",")

	// Declaration of UI Elements
	grid := tview.NewGrid()
	search := tview.NewTextView()
	songs := tview.NewTextView()
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
	songs.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Songs").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	// Setting up UI Elements
	grid.
		SetRows(3, 0, 5).
		SetColumns(30, 0, 20).
		// row index, col index, row span, col span, minheight, minwidth, bool focus
		// Layout for screens narrower than 100 cells (menu and side bar are hidden).
		AddItem(search, 0, 0, 1, 3, 0, 0, false).
		AddItem(nowplaying, 2, 0, 1, 3, 0, 0, false).
		AddItem(songs, 1, 1, 1, 2, 0, 0, false).
		AddItem(playlists, 0, 0, 0, 0, 0, 0, true).
		// Layout for screens wider than 100 cells.
		AddItem(playlists, 1, 0, 1, 1, 0, 100, true)

	// Set up playlist selector
	var sb strings.Builder
	for i, playlist := range playlistArr {
		sb.WriteString(fmt.Sprintf("[\"%d\"]%s[\"\"]\n", i, playlist))
	}


	songs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if songs.HasFocus() {
			switch {
			// Up event
			case event.Rune() == 'j' || event.Key() == tcell.KeyDown:
				// highlightedPl++
			// Down event
			case event.Rune() == 'k' || event.Key() == tcell.KeyUp:
				// highlightedPl--
			case event.Rune() == 'h' || event.Key() == tcell.KeyLeft:
				app.SetFocus(playlists)
				// TODO selection event to show songs in main screen
			case event.Rune() == ' ' || event.Key() == tcell.KeyEnter || event.Rune() == 'l' || event.Key() == tcell.KeyRight:
				// app.SetFocus(songs)
				// TODO selection event to show songs in main screen
			}

			// Go checks modulo by following the sign of a in a%b, python does the opposite
			if highlightedPl < 0 {
				highlightedPl = len(playlistArr) - 1
				playlists.Highlight(fmt.Sprintf("%d", highlightedPl))
			} else {
				playlists.Highlight(fmt.Sprintf("%d", highlightedPl%len(playlistArr)))
			}
		}
		return event
	})

	if err := app.SetRoot(grid, true).SetFocus(playlists).Run(); err != nil {
		panic(err)
	}
}
