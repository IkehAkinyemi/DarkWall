package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// AudioPlayer retrieve music file, decodes and stream the data
func AudioPlayer(musicfile string) {
	// CLI user interface
	err := ui.Init()
	if err != nil {
		log.Fatalf("Failed to initialize UI: %+v", err)
	}
	defer ui.Close()

	file, err := os.Open(musicfile)
	if err != nil {
		log.Fatal(err)
	}

	// decodes audio data in MP3 format to stream the audio.
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	//Initialize speakers
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Setup basic controls for pause/resume music and increase/decrease vol.
	controls := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	vol := &effects.Volume{
		Streamer: controls,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speed := beep.ResampleRatio(4, 1, vol)
	speaker.Play(speed)

	// Extract music file name.
	splitedStr := strings.Split(musicfile, "/")
	selectedMusic := "Current Playing: " + splitedStr[len(splitedStr)-1]

	// CLI user interface: Header content
	headerPara := widgets.NewParagraph()
	headerPara.Title = "DarkWalls🌈"
	headerPara.Text = selectedMusic
	headerPara.TitleStyle.Fg = ui.Color(220)
	headerPara.BorderStyle.Fg = ui.Color(85)
	headerPara.TextStyle.Fg = ui.Color(195)
	headerPara.SetRect(0, 0, 69, 3)

	// CLI user interface: Audio controls content
	ctrlsInfo := widgets.NewParagraph()
	ctrlsInfo.Title = "Audio controls"
	ctrlsInfo.Text = `Pause/resume audio: [SPACE]
	Increase/decrease volume: [↓ ↑] 
	Go back: [ENTER]
	Noraml Speed: [N]
	Speed:  [← →]
	Exit DarkWalls: [Q]`
	ctrlsInfo.TitleStyle.Fg = ui.Color(220)
	ctrlsInfo.BorderStyle.Fg = ui.Color(85)
	ctrlsInfo.TextStyle.Fg = ui.Color(195)
	ctrlsInfo.SetRect(0, 4, 69, 12)

	// Render the UIs
	events := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	render := func() {
		ui.Render(headerPara, ctrlsInfo)
	}

	for {
		select {
		case e := <-events:
			switch e.ID {
			case "<Enter>": // menu
				Start()
			case "<Space>": // pause/resume music
				controls.Paused = !controls.Paused
			case "q": // quit audioplayer
				os.Exit(0)
			case "<Down>": // decrease volume
				vol.Volume -= 0.2
			case "<Up>": //increase volume
				vol.Volume += 0.2
			case "<Left>": // decrease speed by x1.1
				speed.SetRatio(speed.Ratio() - 0.1)
			case "<Right>": // increase speed by x1.1
				speed.SetRatio(speed.Ratio() + 0.1)
			case "n": // Normalize speed
				speed.SetRatio(1)

			}
		case <-ticker:
			render()

		}

		switch {
		case vol.Volume >= 2:
			vol.Volume = 2
		case vol.Volume <= -2:
			vol.Volume = -2
		case speed.Ratio() <= 0.5:
			speed.SetRatio(0.5)
		case speed.Ratio() >= 2:
			speed.SetRatio(2)
		}
	}
}
