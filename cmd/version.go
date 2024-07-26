package cmd

import (
	"bapt/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bapt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bapt version:", config.VERSION)
	},
}
