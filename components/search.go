package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SearchComponent struct{
	searchString string
	searchView *tview.TextView
}

func NewSearchComponent(gui *Gui) *SearchComponent{
	search := tview.NewTextView()
	// Make all elements transparent bg
	search.
		SetBackgroundColor(tcell.ColorNone).
		SetTitle(" Search ").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	return &SearchComponent{
		searchString: "",
		searchView: search,
	}
}
