package cmd

import (
	"strings"

	"github.com/ikehakinyemi/cmd/internal"
	"github.com/spf13/cobra"
)

var setpathCmd = &cobra.Command{
	Use:   "setpath",
	Short: "Create new filepath",
	Long:  `Configure DarkWalls to the filepath (absolute path) that contains the mp3 files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argsString := strings.Join(args, "")

		internal.NewFilePath(argsString)
		internal.AudioMenu()
	},
}

func init() {
	rootCmd.AddCommand(setpathCmd)
}
