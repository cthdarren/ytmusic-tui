package components

import (
	"fmt"

	"github.com/rivo/tview"
)

type Container struct {
	containerView *tview.Grid
}

func NewContainer(gui *Gui) *Container {
	fmt.Println(gui.search.searchView)
	grid := tview.NewGrid()

	grid.
		SetRows(3, 0, 5).
		SetColumns(30, 0, 20).
		// row index, col index, row span, col span, minheight, minwidth, bool focus
		// Layout for screens narrower than 100 cells (menu and side bar are hidden).
		AddItem(gui.search.searchView, 0, 0, 1, 3, 0, 0, false).
		AddItem(gui.nowplaying.nowplayingView, 2, 0, 1, 3, 0, 0, false).
		AddItem(gui.songs.songsView, 1, 1, 1, 2, 0, 0, false).
		AddItem(gui.playlists.playlistsView, 0, 0, 0, 0, 0, 0, true).
		// Layout for screens wider than 100 cells.
		AddItem(gui.playlists.playlistsView, 1, 0, 1, 1, 0, 100, true)

	return &Container{
		containerView: grid,
	}
}
