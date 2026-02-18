package cmd

import (
	"github.com/MohammadTaghipour/flumint/internal/service"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Flumint health and status of components",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunDoctor(cmd)
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunBuild(cmd)
	},
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch between clients",
	RunE: func(cmd *cobra.Command, args []string) error {
		return service.RunCheckout(cmd)
	},
}

func init() {
	buildCmd.Flags().String("client", "", "Client name")
	buildCmd.Flags().String("path", ".", "Path to Flutter project. Default: current directory")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	_ = buildCmd.MarkFlagRequired("client")

	checkoutCmd.Flags().String("client", "", "Client name")
	checkoutCmd.Flags().String("path", ".", "Path to Flutter project. Default: current directory")
	_ = checkoutCmd.MarkFlagRequired("client")
}
