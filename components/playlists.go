package components

import (
	"strings"
	"strconv"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PlaylistComponent struct{
	highlightedPl int;
}


func NewPlaylistComponent(highlightedPl int) *PlaylistComponent{

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
				app.SetFocus(songs)
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
}
