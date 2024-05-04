package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SongsComponent struct{
	// probably an array/slice/map of songs
	songData []int
	highlightedSong int
	songsView tview.TextView
}

func NewSongsComponent(gui *Gui) *SongsComponent{
	songs := tview.NewTextView()
	var songData []int
	highlighted := 0
	songs.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Songs").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

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
				gui.app.SetFocus(&gui.playlists.playlistsView)
				// TODO selection event to show songs in main screen
			case event.Rune() == ' ' || event.Key() == tcell.KeyEnter || event.Rune() == 'l' || event.Key() == tcell.KeyRight:
				// app.SetFocus(songs)
				// TODO selection event to show songs in main screen
			}

			// Go checks modulo by following the sign of a in a%b, python does the opposite
			if gui.playlists.highlightedPl < 0 {
				highlighted= len(songData) - 1
				songs.Highlight(fmt.Sprintf("%d", highlighted))
			} else {
				songs.Highlight(fmt.Sprintf("%d", highlighted%len(songData)))
			}
		}
		return event
	})

	return &SongsComponent{
		songData: songData,
		highlightedSong: highlighted,
		songsView: *tview.NewTextView(),
	}
}
