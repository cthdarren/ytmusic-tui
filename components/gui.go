package components

import "github.com/rivo/tview"

type State struct{
	currPlaylistIndex int;
}

type Gui struct {
	app *tview.Application
	state *State
}
