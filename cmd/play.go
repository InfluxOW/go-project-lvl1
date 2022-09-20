package cmd

import (
	"github.com/InfluxOW/go-project-lvl1/internal/app"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game",
	Run: func(cmd *cobra.Command, args []string) {
		app.Play(random)
	},
}

var (
	random = false
)

func init() {
	rootCmd.AddCommand(playCmd)

	playCmd.Flags().BoolVar(&random, "random", false, "play random game")
}
