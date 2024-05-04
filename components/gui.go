package components

import "github.com/rivo/tview"

type State struct{

}

type Gui struct {
	app *tview.Application
	playlists *PlaylistComponent
	songs *SongsComponent
	search *SearchComponent
	nowplaying *NowPlayingComponent
	state *State
}

func NewState() *State{
	return &State{
	}
}

func NewGui() *Gui{
	return &Gui{
		app: tview.NewApplication(),
		playlists: NewPlaylistComponent(),
		songs: NewSongsComponent(),
		search: NewSearchComponent(),
		nowplaying: NewNowPlayingComponent(),
		state: NewState(),
	}
}
