package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd is the main command for Flumint CLI.
var rootCmd = &cobra.Command{
	Use:   "flumint",
	Short: "Flumint - Multi-client Flutter Build Tool",
	Long:  "Flumint is a CLI tool to manage, build, and switch between multiple Flutter clients efficiently.",
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(checkoutCmd)
}
