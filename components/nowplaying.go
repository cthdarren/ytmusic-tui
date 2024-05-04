package components

import "github.com/rivo/tview"

type NowPlayingComponent struct{
	currentSong int
	shuffle bool
	loop bool
	nowplayingView tview.TextView
}
func NewNowPlayingComponent(gui *Gui) *NowPlayingComponent{
	return &NowPlayingComponent{
		currentSong: 0,
		shuffle: false,
		loop: false,
		nowplayingView: *tview.NewTextView(),
	}
}
