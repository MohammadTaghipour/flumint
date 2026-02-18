package cmd

import (
	"github.com/MohammadTaghipour/flumint/internal/service"
	"github.com/spf13/cobra"
)

// networkCmd checks network connectivity to all repositories needed for Flutter development.
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Check network connectivity to required repositories",
	Long: `Checks access to all major repositories required for Flutter development:

- pub.dev
- Flutter Storage
- Google Maven
- Maven Central
- CocoaPods CDN

Displays the connection status for each repository along with latency.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunNetworkCheck()
	},
}

// doctorCmd shows Flumint health and version information.
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Flumint health and status of components",
	Long:  "Displays Flutter version, Dart version, DevTools version, and other relevant system information for Flumint.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunDoctor(cmd)
	},
}

// buildCmd builds a Flutter project for a specific client.
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	Long:  "Builds a Flutter project for the specified client, platform, and environment, injecting assets and updating configuration automatically.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunBuild(cmd)
	},
}

// checkoutCmd switches the project configuration to another client.
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch between clients",
	Long:  "Updates the Flutter project to use another client's configuration, including app name, package name, and assets.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunCheckout(cmd)
	},
}

func init() {
	// Build command flags
	buildCmd.Flags().String("client", "", "Client name (required)")
	buildCmd.Flags().String("path", ".", "Path to Flutter project (default: current directory)")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	_ = buildCmd.MarkFlagRequired("client")

	// Checkout command flags
	checkoutCmd.Flags().String("client", "", "Client name (required)")
	checkoutCmd.Flags().String("path", ".", "Path to Flutter project (default: current directory)")
	_ = checkoutCmd.MarkFlagRequired("client")
}
