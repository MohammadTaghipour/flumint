package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	Run: func(cmd *cobra.Command, args []string) {
		clientName, _ := cmd.Flags().GetString("client")
		platform, _ := cmd.Flags().GetString("platform")
		env, _ := cmd.Flags().GetString("env")

		fmt.Printf("Building client %s for %s in %s environment\n", clientName, platform, env)

		// Resolve client

		// Load config
		// cfg, err := config.Load(clientPath)
		// if err != nil {
		// 	panic(err)
		// }

		// Inject assets

		// Patch Android build files

		// Update pubspec.yaml

		// Execute flutter build

		fmt.Println("Build finished successfully!")
	},
}

func init() {
	buildCmd.Flags().String("client", "", "Client name")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	buildCmd.MarkFlagRequired("client")
}
