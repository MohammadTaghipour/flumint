package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MohammadTaghipour/flumint/internal/assets"
	"github.com/MohammadTaghipour/flumint/internal/client"
	"github.com/MohammadTaghipour/flumint/internal/config"
	"github.com/MohammadTaghipour/flumint/internal/flutter"
	"github.com/MohammadTaghipour/flumint/internal/platform/android"
	"github.com/MohammadTaghipour/flumint/internal/platform/web"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Flumint health and status of components",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Running Flumint doctor...\n")

		flutterV, err := flutter.GetVersion()
		if err != nil {
			return fmt.Errorf("failed to get flutter version. %w", err)
		}

		fmt.Printf("Flutter version: %s\n", flutterV.Version)
		fmt.Printf("Flutter channel: %s\n", flutterV.Channel)
		fmt.Printf("Dart version: %s\n", flutterV.Dart)
		fmt.Printf("DevTools: %s\n", flutterV.DevTools)

		// TODO: uncomment if needed
		// flutterD, err := flutter.RunDoctor()
		// if err != nil {
		// 	return fmt.Errorf("failed run flutter doctor. %w", err)
		// }
		// fmt.Println(flutterD)

		return nil
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := os.Getwd()
		if err != nil {
			return err
		}

		clientName, _ := cmd.Flags().GetString("client")
		platform, _ := cmd.Flags().GetString("platform")
		env, _ := cmd.Flags().GetString("env")
		projectRoot, _ := cmd.Flags().GetString("path")

		if projectRoot != "" {
			root = projectRoot
		}

		fmt.Printf("Building client %s for %s in %s environment\n", clientName, platform, env)

		// Resolve client
		clientPath, err := client.Resolve(root, clientName)
		if err != nil {
			return fmt.Errorf("failed to resolve client: %w", err)
		}

		// Load config
		cfg, err := config.Load(filepath.Join(root, clientName))
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Inject assets
		if err := assets.Inject(root, clientPath); err != nil {
			return fmt.Errorf("failed to inject assets: %w", err)
		}

		switch platform {
		case "android":
			androidUtil := android.NewAndroid(root)

			// oldAppName, err := androidUtil.GetAppName()
			// if err != nil {
			// 	return fmt.Errorf("failed to fetch android app name: %w", err)
			// }

			// if oldAppName != cfg.AppName {
			// 	if err := androidUtil.SetAppName(cfg.AppName); err != nil {
			// 		return fmt.Errorf("failed to set android app name: %w", err)
			// 	}
			// }

			oldPackageName, err := androidUtil.GetPackageName()
			if err != nil {
				return fmt.Errorf("failed to fetch android package name: %w", err)
			} else {
				fmt.Printf("Old package name is %s\n", oldPackageName)
			}
			// if oldPackageName != cfg.PackageName {
			// 	if err := androidUtil.SetBundleId(cfg.PackageName); err != nil {
			// 		return fmt.Errorf("failed to set android package name: %w", err)
			// 	}
			// }
		case "web":
			webUtil := web.NewWeb("./")

			oldAppName, err := webUtil.GetAppName()
			if err != nil {
				return fmt.Errorf("failed to fetch web app name: %w", err)
			}
			if oldAppName != cfg.AppName {
				if err := webUtil.SetAppName(cfg.AppName); err != nil {
					return fmt.Errorf("failed to set web app name: %w", err)
				}
			}
		default:
			return fmt.Errorf("unsupported platform: %s", platform)
		}

		// Update pubspec.yaml

		// Execute flutter build
		if err := flutter.Build(platform, clientName, cfg); err != nil {
			return fmt.Errorf("failed to build app. error: %w", err)
		}

		fmt.Println("Build finished successfully!")

		return nil
	},
}

func init() {
	buildCmd.Flags().String("client", "", "Client name")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	_ = buildCmd.MarkFlagRequired("client")
}
