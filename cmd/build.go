package cmd

import (
	"flumint/internal/assets"
	"flumint/internal/client"
	"flumint/internal/config"
	"flumint/internal/flutter"
	"flumint/internal/platform/android"
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
		clientPath, err := client.Resolve(clientName)
		if err != nil {
			panic(err)
		}

		// Load config
		cfg, err := config.Load(clientPath)
		if err != nil {
			panic(err)
		}
		println(cfg.PackageName)

		// Inject assets
		if err := assets.Inject(clientPath); err != nil {
			panic(err)
		}

		// Patch Android build files
		if platform == "android" {
			androidUtil, err := android.NewAndroid("./")
			if err != nil {
				panic(err)
			}

			oldAppName, err := androidUtil.GetAppName()
			if err != nil {
				panic(err)
			}
			if oldAppName != cfg.AppName {
				if err := androidUtil.SetAppName(cfg.AppName); err != nil {
					panic(err)
				}
			}

			oldPackageName, err := androidUtil.GetBundleId()
			if err != nil {
				panic(err)
			}
			if oldPackageName != cfg.PackageName {
				if err := androidUtil.SetBundleId(cfg.PackageName); err != nil {
					panic(err)
				}
			}
		}

		// Update pubspec.yaml

		// Execute flutter build
		if err := flutter.Build(platform, clientName, cfg); err != nil {
			panic(err)
		}

		fmt.Println("Build finished successfully!")
	},
}

func init() {
	buildCmd.Flags().String("client", "", "Client name")
	buildCmd.Flags().String("platform", "android", "Target platform: android/web")
	buildCmd.Flags().String("env", "dev", "Environment: dev/staging/prod")
	buildCmd.MarkFlagRequired("client")
}
