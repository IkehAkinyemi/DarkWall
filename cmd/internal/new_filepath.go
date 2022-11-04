package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewFilePath creates a new music folder if none is found
func NewFilePath(newPath string) {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("music_path")
	v.SetConfigType("json")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// write to config.json file the absolute path for your music folder
	v.Set("music_path", newPath)
	v.WriteConfig()

	fmt.Println("Folder successfully added")

}
