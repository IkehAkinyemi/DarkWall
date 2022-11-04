package internal

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/viper"
)

// AudioMenu setups configs and the AudioPlayer UI
func AudioMenu() {
	v := viper.New()

	// Retrieve music path stored in music_path.json
	v.AddConfigPath(".")
	v.SetConfigName("./music_path")
	v.SetConfigType("json")

	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	// Retrieve directory entries
	filepath, err := os.ReadDir(v.GetString("filepath"))
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
	list.Title = "Your music list"
	list.Rows = files
	list.TextStyle = ui.NewStyle(ui.Color(195))
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
			AudioPlayer(fmt.Sprintf("%s/%s", v.GetString("filepath"), selectedFIle))
		}
		ui.Render(list)
	}
}
