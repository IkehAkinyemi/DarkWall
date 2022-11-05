package internal

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)
type AudioPlayer struct {
	Directory string
}

// AudioMenu setups configs and the AudioPlayer UI
func (au AudioPlayer) AudioMenu() {

	// Retrieve directory entries
	filepath, err := os.ReadDir(au.Directory)
	if err != nil {
		log.Fatal(err)
	}

	var files []string

	for _, file := range filepath {
		files = append(files, file.Name())
	}

	err = ui.Init()
	if err != nil {
		log.Fatalf("Failed to load termui: %+v", err)
	}
	defer ui.Close()

	// CLI user interface: Music list content
	list := widgets.NewList()
	list.Title = "DarkWalls🌈: Select music file to play"
	list.Rows = files
	list.SelectedRowStyle = ui.NewStyle(ui.Color(207))
	list.WrapText = false
	list.SetRect(0, 0, 69, len(files)+2)
	list.BorderStyle.Fg = ui.Color(85)

	ui.Render(list)

	prevKey := ""
	events := ui.PollEvents()

	for {
		event := <-events
		switch event.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		case "<C-d>":
			list.ScrollHalfPageDown()
		case "<C-u>":
			list.ScrollHalfPageUp()
		case "<C-f>":
			list.ScrollPageDown()
		case "<C-b>":
			list.ScrollPageUp()
		case "g":
			if prevKey == "g" {
				list.ScrollTop()
			}
		case "<Enter>":
			selectedFIle := files[list.SelectedRow]

			// AudioPlayer contains logic for music player
			au.Player(fmt.Sprintf("%s/%s", au.Directory, selectedFIle))
		}
		ui.Render(list)
	}
}
