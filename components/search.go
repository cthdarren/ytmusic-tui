package components

import "github.com/rivo/tview"

type SearchComponent struct{
	searchString string
	searchView tview.InputField
}

func NewSearchComponent(gui *Gui) *SearchComponent{
	return &SearchComponent{
		searchString: "",
	}
}
