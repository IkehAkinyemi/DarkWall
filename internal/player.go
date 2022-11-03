package internal

import (
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

	finished := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(
		func ()  {
			finished <- true
		})))

	<-finished
}