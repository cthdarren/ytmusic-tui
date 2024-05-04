package components

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SongsComponent struct{
	// probably an array/slice/map of songs
	songData []string
	highlightedSong int
	songsView *tview.TextView
}

func NewSongsComponent(gui *Gui) *SongsComponent{
	songs := tview.NewTextView()
	songData := []string{"Hello - Adele", "Deep down - Aimer", "Like a - yama", "Kagirinaku haiiro he - Nightcord at 25:00", "more than words - Hitsujibungaku", "New Genesis (UTA from ONEPIECE FILM RED) - Ado", "Voices of the Chord - Hiroyuki Sawano"} 
	highlighted := 0

	var sb strings.Builder
	for i, song := range songData{
		sb.WriteString(fmt.Sprintf("[\"%d\"]%s[\"\"]\n", i, song))
	}
	songs.
		SetText(sb.String()).
		SetRegions(true).
		Highlight(fmt.Sprintf("%d", highlighted)).
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Songs").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	songs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if songs.HasFocus() {
			switch {
			// Up event
			case event.Rune() == 'j' || event.Key() == tcell.KeyDown:
				highlighted++
			// Down event
			case event.Rune() == 'k' || event.Key() == tcell.KeyUp:
				highlighted--
			case event.Rune() == 'h' || event.Key() == tcell.KeyLeft:
				songs.SetBorderColor(tcell.ColorWhite)
				songs.SetTitleColor(tcell.ColorWhite)
				gui.playlists.playlistsView.SetBorderColor(tcell.ColorRed)
				gui.playlists.playlistsView.SetTitleColor(tcell.ColorRed)
				gui.app.SetFocus(gui.playlists.playlistsView)
				// TODO selection event to show songs in main screen
			case event.Rune() == ' ' || event.Key() == tcell.KeyEnter || event.Rune() == 'l' || event.Key() == tcell.KeyRight:
				// TODO selection event to show songs in main screen
			}

			// Go checks modulo by following the sign of a in a%b, python does the opposite
			if highlighted < 0 {
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
		songsView: songs,
	}
}
