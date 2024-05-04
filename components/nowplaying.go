package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NowPlayingComponent struct{
	currentSong int
	shuffle string 
	loop string
	nowplayingView *tview.TextView
}
func NewNowPlayingComponent(gui *Gui) *NowPlayingComponent{
	shuffleStr := "Off"
	loopStr := "On"
	nowplaying := tview.NewTextView().SetText("Fireworks Festival\nRadwimps")
	nowplaying.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle(fmt.Sprintf(" Playing | Shuffle: %s | Repeat: %s ", shuffleStr, loopStr)).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	return &NowPlayingComponent{
		currentSong: 0,
		shuffle: shuffleStr,
		loop: loopStr,
		nowplayingView: nowplaying,
	}
}
