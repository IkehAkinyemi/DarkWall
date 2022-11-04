package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	termbx "github.com/nsf/termbox-go"
)

// player retrieve music file, decodes and stream the data
func Player(musicfile string) {
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
	vol := effects.Volume {
		Streamer: controls,
		Base: 2,
		Volume: 0,
		Silent: false,
	}
	speaker.Play(controls)

	termboxErr := termbx.Init()
	if termboxErr != nil {
		log.Fatal(termboxErr)
	}
	defer termbx.Close()

	fmt.Printf("Playing: %s\n\n", strings.Replace(musicfile, ".mp3", "", 1))
	fmt.Println("Use [ENTER] to pause/resume: [ENTER]")
	fmt.Println("Increase/decrease volume: [↓ ↑]")
	fmt.Println("Go back: [BACKSPACE]")
	fmt.Println("Press escape key to exit DarkWalls: [Esc]")

	for {
		keyEvent := termbx.PollEvent()

		speaker.Lock()

		switch {
		case keyEvent.Key == termbx.KeyEnter:
			Start()
		case keyEvent.Key == termbx.KeySpace:
			controls.Paused = !controls.Paused
		case keyEvent.Key == termbx.KeyEsc:
			os.Exit(0)
		case keyEvent.Key == termbx.KeyArrowDown:
			vol.Volume -= 0.2
		case keyEvent.Key == termbx.KeyArrowUp:
			vol.Volume += 0.2
		}

		switch {
		case vol.Volume >= 2:
			vol.Volume = 2
		case vol.Volume <= -2:
			vol.Volume = -2
		}

		speaker.Unlock()
	}
}
