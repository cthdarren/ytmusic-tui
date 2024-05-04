package components

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PlaylistComponent struct{
	// TODO playlistData map/array/slice
	playlistData string;
	highlightedPl int;
	playlistsView *tview.TextView
}


func NewPlaylistComponent(gui *Gui) *PlaylistComponent{
	highlightedPl := 0

	// TODO change to actual data
	playliststringtest := "wee duty mood,chill,poggers,kpop,jap weirdge songs,weeb,normie,rock,linkin park,nc25,bandori"
	playlistArr := strings.Split(playliststringtest, ",")

	var sb strings.Builder
	playlists := tview.NewTextView()
	playlists.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Playlists").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	// Key bindings for highlighting of playlists for selection
	playlists.SetRegions(true).SetText(sb.String())
	if playlists.HasFocus() {
		playlists.Highlight(strconv.Itoa(highlightedPl))
	}

	// Set up playlist selector
	for i, playlist := range playlistArr {
		sb.WriteString(fmt.Sprintf("[\"%d\"]%s[\"\"]\n", i, playlist))
	}

	playlists.SetText(sb.String())
	playlists.Highlight(fmt.Sprintf("%d", highlightedPl))

	playlists.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if playlists.HasFocus() {
			switch {
			// Up event
			case event.Rune() == 'j' || event.Key() == tcell.KeyDown:
				highlightedPl++
			// Down event
			case event.Rune() == 'k' || event.Key() == tcell.KeyUp:
				highlightedPl--
			case event.Rune() == ' ' || event.Key() == tcell.KeyEnter || event.Rune() == 'l' || event.Key() == tcell.KeyRight:
				gui.app.SetFocus(gui.songs.songsView)
				playlists.SetBorderColor(tcell.ColorWhite)
				playlists.SetTitleColor(tcell.ColorWhite)
				gui.songs.songsView.SetBorderColor(tcell.ColorRed)
				gui.songs.songsView.SetTitleColor(tcell.ColorRed)
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

	return &PlaylistComponent{
		highlightedPl: 0,
		playlistsView: playlists,
	}
}
