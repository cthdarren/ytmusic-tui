package main

import (
	"fmt"
	"github.com/rivo/tview"
)

// User Interface
func main() {

	artistsBox := tview.NewTextView()
	w := artistsBox.BatchWriter();
	defer artistsBox.Clear();
	fmt.Fprintln(w, "Nightcord at 25:00");
	fmt.Fprintln(w, "Roselia");
	fmt.Fprintln(w, "Linkin Park");
	fmt.Fprintln(w, "IVE");
	fmt.Fprintln(w, "Le Sserafim");
	fmt.Println(artistsBox.GetText(false))

	//box := tview.NewBox().SetBorder(true).SetTitle("YTM TUI!");

	if err := tview.NewApplication().SetRoot(artistsBox, true).Run(); err != nil {
		panic(err)
	}
}
