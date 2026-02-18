package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flumint",
	Short: "Flumint - Multi-client Flutter Build Tool",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(checkoutCmd)
}
