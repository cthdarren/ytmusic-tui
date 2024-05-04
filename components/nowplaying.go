package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NowPlayingComponent struct{
	currentSong int
	shuffle bool
	loop bool
	nowplayingView *tview.TextView
}
func NewNowPlayingComponent(gui *Gui) *NowPlayingComponent{
	nowplaying := tview.NewTextView().SetText("Fireworks Festival\nRadwimps")
	nowplaying.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle("Playing").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	return &NowPlayingComponent{
		currentSong: 0,
		shuffle: false,
		loop: false,
		nowplayingView: nowplaying,
	}
}
