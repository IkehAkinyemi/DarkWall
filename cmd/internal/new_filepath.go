package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//NewFilePath creates a new music folder if none is found
func NewFilePath() {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("music_path")
	v.SetConfigType("json")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Only ask and retrieve folder containing music when none is found
	if v.Get("music_path") == "" {
		fmt.Println("Enter folder (absolute path) containing music files")
		scn := bufio.NewScanner(os.Stdin)
		scn.Scan()

		newPath := scn.Text()

		v.Set("music_path", newPath)
		v.WriteConfig()

		fmt.Println("Folder successfully added")
	}
}