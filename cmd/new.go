package cmd

import (
	"github.com/ikehakinyemi/cmd/internal"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new filepath",
	Long: `Configure DarkWalls to the filepath (absolute path) that contains the mp3 files`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.NewFilePath()
		internal.Start()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}