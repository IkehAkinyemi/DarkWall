package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ikehakinyemi/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func main() {
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

	// Setup a select dialog for user
	prompt := promptui.Select {
		Label: "Select music file to play:",
		Items: files,
	}
	_, musicFile, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt to select music failed: %+v \n", err)
	}

	internal.Player(fmt.Sprintf("%s/%s", v.GetString("filepath"), musicFile))	
}