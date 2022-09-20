package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-project-lvl1",
	Short: "A pack of simple math games.",
}

func Execute() {
	rootCmd.Execute()
}
