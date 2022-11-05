package main

import (
	"flag"

	"github.com/ikehakinyemi/dark-walls/internal"
)

func main() {
	var audioPlayer internal.AudioPlayer

	flag.StringVar(&audioPlayer.Directory, "music-dir", "/Users/ikehakinyemi/Projects/Go/DarkWalls🌈/sampleMusic", "Specify the absolute path to your music directory")
	flag.Parse()

	audioPlayer.AudioMenu()
}
