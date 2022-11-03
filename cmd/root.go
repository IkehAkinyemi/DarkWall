package cmd

import (
	"os"

	"github.com/ikehakinyemi/cmd/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "DarkWall",
	Short: "DarkWall music player for the CLI",
	Long: "DarkWalls🌈 is a music player of the CLI, by the CLI and for the CLI (pun intended)",
	Run: func (cmd *cobra.Command, args []string)  {
		internal.Start()
	},
}

func Initiate() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}