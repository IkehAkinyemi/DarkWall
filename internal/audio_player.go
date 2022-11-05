package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// AudioPlayer retrieve music file, decodes and stream the data
func (au AudioPlayer) Player(musicfile string) {
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
	defer file.Close()

	// decodes audio data in MP3 format to stream the audio.
	streamer, format, err := fileFormat(musicfile, file)
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
	defer speaker.Close()
	

	// Extract music file name.
	splitedStr := strings.Split(musicfile, "/")
	selectedMusic := "Current Playing: " + splitedStr[len(splitedStr)-1]

	// CLI UI: Header content
	headerPara := widgets.NewParagraph()
	headerPara.Title = "DarkWalls🌈"
	headerPara.Text = selectedMusic
	headerPara.TitleStyle.Fg = ui.Color(220)
	headerPara.BorderStyle.Fg = ui.Color(85)
	headerPara.TextStyle.Fg = ui.Color(195)
	headerPara.SetRect(0, 0, 69, 3)

	// CLI UI: Audio controls content
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
	ctrlsInfo.SetRect(0, 9, 69, 16)

	// CLI UI: speed rate UI
	speedGauge := widgets.NewGauge()
	speedGauge.Title = "Speed rate"
	speedGauge.Percent = 50
	speedGauge.BarColor = ui.Color(185)
	speedGauge.BorderStyle.Fg = ui.Color(85)
	speedGauge.TitleStyle.Fg = ui.Color(220)
	speedGauge.SetRect(0, 6, 69, 9)

	// CLI UI: volume UI
	volGuage := widgets.NewGauge()
	volGuage.Title = "Volume"
	volGuage.Percent = 50
	volGuage.BarColor = ui.ColorMagenta
	volGuage.BorderStyle.Fg = ui.Color(85)
	volGuage.TitleStyle.Fg = ui.Color(220)
	volGuage.SetRect(0, 3, 69, 6)

	// Render the UIs
	events := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	render := func() {
		ui.Render(headerPara, ctrlsInfo, speedGauge, volGuage)
	}

	for {
		select {
		case e := <-events:
			switch e.ID {
			case "<Enter>": // menu
				controls.Paused = !controls.Paused
				streamer.Close()
				speaker.Close()
				ui.Close()
				file.Close()

				// Return to menu
				os.Exit(0)
			case "<Space>": // pause/resume music
				controls.Paused = !controls.Paused
			case "q": // quit audioplayer
				os.Exit(0)
			case "<Down>": // decrease volume
				volGuage.Percent -= 2
				vol.Volume -= 0.2
			case "<Up>": //increase volume
				volGuage.Percent += 2
				vol.Volume += 0.2
			case "<Left>": // decrease speed by x1.1
				speedGauge.Percent -= 2
				speed.SetRatio(speed.Ratio() - 0.1)
			case "<Right>": // increase speed by x1.1
				speedGauge.Percent += 2
				speed.SetRatio(speed.Ratio() + 0.1)
			case "n": // Normalize speed
				speed.SetRatio(1)

			}
		case <-ticker:
			render()

		}

		// Revert to safe threshold
		switch {
		case vol.Volume >= 2:
			volGuage.Percent = 100
			vol.Volume = 2
		case vol.Volume <= -2:
			vol.Volume = -2
			volGuage.Percent = 0
		case speed.Ratio() <= 0.5:
			speedGauge.Percent = 0
			speed.SetRatio(0.5)
		case speed.Ratio() >= 2:
			speedGauge.Percent = 100
			speed.SetRatio(2)
		}
	}
}

// fileFormat checks file extension and returen right decoder for it.
func fileFormat(fileName string, file *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	switch {
	case strings.HasSuffix(fileName, ".wav"):
		return wav.Decode(file)
	case strings.HasSuffix(fileName, ".flac"):
		return flac.Decode(file)
	case strings.HasSuffix(fileName, ".ogg"):
		return vorbis.Decode(file)
	default:
		return mp3.Decode(file)
	} 
}