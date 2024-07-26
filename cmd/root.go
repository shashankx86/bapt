package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bapt",
	Short: "bapt: An experimental package manager",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands to the root command here
	rootCmd.AddCommand(versionCmd)
}
