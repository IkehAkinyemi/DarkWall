package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// player retrieve music file, decodes and stream the data
func Player(musicPath string) {
	file, err := os.Open(musicPath)
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

	controls := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	speaker.Play(controls)

	log.Print("Press [ENTER] to pause/resume")

	for {
		fmt.Scanln()
		speaker.Lock()
		controls.Paused = !controls.Paused
		speaker.Unlock()
	}
}