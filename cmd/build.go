package cmd

import (
	"fmt"

	"github.com/MohammadTaghipour/flumint/internal/assets"
	"github.com/MohammadTaghipour/flumint/internal/client"
	"github.com/MohammadTaghipour/flumint/internal/config"
	"github.com/MohammadTaghipour/flumint/internal/flutter"
	"github.com/MohammadTaghipour/flumint/internal/platform/android"
	"github.com/MohammadTaghipour/flumint/internal/platform/web"
	"github.com/MohammadTaghipour/flumint/internal/utils"
	"github.com/briandowns/spinner"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Flumint health and status of components",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println()

		s := spinner.New(spinner.CharSets[utils.SpinnerCharset], utils.SpinnerDuration)
		s.Suffix = " Running Flumint doctor..."
		s.Color(utils.SpinnerColor)
		s.Start()

		flutterV, err := flutter.GetVersion()
		if err != nil {
			s.Stop()
			fmt.Println("Flumint doctor failed ✖")
			return fmt.Errorf("failed to get flutter version: %w", err)
		}

		s.Stop()
		fmt.Println(utils.BrandWriter("Flumint Doctor Information"))
		fmt.Println("------------------------")
		fmt.Printf("Version   : %s\n", flutterV.Version)
		fmt.Printf("Channel   : %s\n", flutterV.Channel)
		fmt.Printf("Dart      : %s\n", flutterV.Dart)
		fmt.Printf("DevTools  : %s\n", flutterV.DevTools)
		fmt.Println()

		return nil
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	RunE: func(cmd *cobra.Command, args []string) error {

		clientName, _ := cmd.Flags().GetString("client")
		platform, _ := cmd.Flags().GetString("platform")
		env, _ := cmd.Flags().GetString("env")
		root, _ := cmd.Flags().GetString("path")

		fmt.Printf("Building client %s for %s in %s environment\n", clientName, platform, env)

		// Resolve client
		clientPath, err := client.Resolve(root, clientName)
		if err != nil {
			return fmt.Errorf("failed to resolve client: %w", err)
		}

		// Load config
		cfg, err := config.Load(clientPath)
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

			oldAppName, err := androidUtil.GetAppName()
			if err != nil {
				return fmt.Errorf("failed to fetch android app name: %w", err)
			}

			if oldAppName != cfg.AppName {
				if err := androidUtil.SetAppName(cfg.AppName); err != nil {
					return fmt.Errorf("failed to set android app name: %w", err)
				}
			}

			oldPackageName, err := androidUtil.GetPackageName()
			if err != nil {
				return fmt.Errorf("failed to fetch android package name: %w", err)
			} else {
				fmt.Printf("Old package name is %s\n", oldPackageName)
			}

			if oldPackageName != cfg.PackageName {
				if err := androidUtil.SetPackageName(cfg.PackageName); err != nil {
					return fmt.Errorf("failed to set android package name: %w", err)
				}
			}

			if err := androidUtil.SetPackageNameInActivities(cfg.PackageName); err != nil {
				return fmt.Errorf("failed to set android package name in java|kotlin files: %w", err)
			}

			if err := androidUtil.SetPackageNameInManifest(cfg.PackageName); err != nil {
				return fmt.Errorf("failed to set android package name in manifest files: %w", err)
			}

		case "web":
			webUtil := web.NewWeb(root)

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
		// if err := flutter.Build(root, platform, clientName, cfg); err != nil {
		// 	return fmt.Errorf("failed to build app. error: %w", err)
		// }

		fmt.Println("Build finished successfully!")

		return nil
	},
}

func init() {
	buildCmd.Flags().String("client", "", "Client name")
	buildCmd.Flags().String("path", ".", "Path to Flutter project. Default: current root")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	_ = buildCmd.MarkFlagRequired("client")
}
