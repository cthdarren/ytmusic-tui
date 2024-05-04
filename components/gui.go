package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type State struct{

}

type Gui struct {
	app *tview.Application
	container *Container
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
		state: NewState(),
	}
}

func (g *Gui) init(){
	g.playlists = NewPlaylistComponent(g)
	g.songs = NewSongsComponent(g)
	g.search = NewSearchComponent(g)
	g.nowplaying = NewNowPlayingComponent(g)
	g.container = NewContainer(g)
}

func (g *Gui) Start() error{
	g.init()
	g.playlists.playlistsView.SetBorderColor(tcell.ColorRed)
	g.playlists.playlistsView.SetTitleColor(tcell.ColorRed)
	if err := g.app.SetRoot(g.container.containerView, true).SetFocus(g.playlists.playlistsView).Run(); err != nil{
		g.app.Stop()
		return err
	}
	return nil
}

func (g *Gui) Goto(from *tview.Box, to *tview.Box){
	from.SetBorderColor(tcell.ColorWhite)
	from.SetTitleColor(tcell.ColorWhite)
	to.SetBorderColor(tcell.ColorRed)
	to.SetTitleColor(tcell.ColorRed)
	g.app.SetFocus(to)
}



