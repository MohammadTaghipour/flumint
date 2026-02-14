package cmd

import (
	"fmt"
	"time"

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
		fmt.Println(utils.SuccessWriter("Version    : " + flutterV.Version))
		fmt.Println(utils.SuccessWriter("Channel    : " + flutterV.Channel))
		fmt.Println(utils.SuccessWriter("Dart       : " + flutterV.Dart))
		fmt.Println(utils.SuccessWriter("DevTools   : " + flutterV.DevTools))

		fmt.Println()

		return nil
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Flutter project for a client",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println()

		s := spinner.New(spinner.CharSets[utils.SpinnerCharset], utils.SpinnerDuration)
		s.Suffix = " Running Flumint build..."
		s.Color(utils.SpinnerColor)
		s.Start()

		time.Sleep(time.Second)

		s.Stop()

		clientName, _ := cmd.Flags().GetString("client")
		platform, _ := cmd.Flags().GetString("platform")
		env, _ := cmd.Flags().GetString("env")
		root, _ := cmd.Flags().GetString("path")

		fmt.Printf("Building client %s for %s in %s environment...\n", clientName, platform, env)

		clientPath, err := client.Resolve(root, clientName)
		if err != nil || clientPath == "" {
			fmt.Println(utils.ErrorWriter("Build failed."))
			return fmt.Errorf("failed to resolve client: %w", err)
		}
		fmt.Printf("client path resolved: %s\n", clientPath)

		cfg, err := config.Load(clientPath)
		if err != nil {
			fmt.Println(utils.ErrorWriter("Build failed."))
			return fmt.Errorf("failed to load config: %w", err)
		}
		fmt.Println("config.json detected")

		fmt.Println(utils.InfoWriter(fmt.Sprintf("AppName     : %s", cfg.AppName)))
		fmt.Println(utils.InfoWriter(fmt.Sprintf("PackageName : %s", cfg.PackageName)))

		s.Suffix = " Injecting client assets and files..."
		s.Start()
		time.Sleep(time.Second * 2)
		if err := assets.Inject(root, clientPath); err != nil {
			fmt.Println(utils.ErrorWriter("Build failed."))
			return fmt.Errorf("failed to inject assets: %w", err)
		}
		s.Stop()
		fmt.Println("Assets injected sucessfully")

		switch platform {
		case "android":
			androidUtil := android.NewAndroid(root)

			currAppName, err := androidUtil.GetAppName()
			if err != nil {
				fmt.Println(utils.ErrorWriter("Build failed."))
				return fmt.Errorf("failed to fetch android app name: %w", err)
			}
			fmt.Printf("current app name is: %s\n", utils.InfoWriter(currAppName))

			if currAppName != cfg.AppName {
				if err := androidUtil.SetAppName(cfg.AppName); err != nil {
					fmt.Println(utils.ErrorWriter("Build failed."))
					return fmt.Errorf("failed to set android app name: %w", err)
				}
				fmt.Printf("app name changed from %s to %s\n", utils.InfoWriter(currAppName), utils.InfoWriter(cfg.AppName))
			} else {
				fmt.Println("no changes in app name. current and new app names are same. ")
			}

			currPackageName, err := androidUtil.GetPackageName()
			if err != nil {
				fmt.Println(utils.ErrorWriter("Build failed."))
				return fmt.Errorf("failed to fetch android package name: %w", err)
			}
			fmt.Printf("current package name is: %s\n", utils.InfoWriter(currPackageName))

			if currPackageName != cfg.PackageName {
				if err := androidUtil.SetPackageName(cfg.PackageName); err != nil {
					fmt.Println(utils.ErrorWriter("Build failed."))
					return fmt.Errorf("failed to set android package name: %w", err)
				}
				fmt.Printf("package name changed from %s to %s\n", utils.InfoWriter(currPackageName), utils.InfoWriter(cfg.PackageName))

				if err := androidUtil.SetPackageNameInActivities(cfg.PackageName); err != nil {
					fmt.Println(utils.ErrorWriter("Build failed."))
					return fmt.Errorf("failed to update package name in java/kotlin files: %w", err)
				}
				fmt.Printf("package name in java|kotlin files changed to %s\n", utils.InfoWriter(cfg.PackageName))

				if err := androidUtil.SetPackageNameInManifest(cfg.PackageName); err != nil {
					fmt.Println(utils.ErrorWriter("Build failed."))
					return fmt.Errorf("failed to update package name in manifest files: %w", err)
				}
				fmt.Printf("package name in AndroidManifest.xml file changed to %s\n", utils.InfoWriter(cfg.PackageName))
			} else {
				fmt.Println("no changes in package name. current and new package names are same. ")
			}

		case "web":
			webUtil := web.NewWeb(root)

			oldAppName, err := webUtil.GetAppName()
			if err != nil {
				fmt.Println(utils.ErrorWriter("Build failed."))
				return fmt.Errorf("failed to fetch web app name: %w", err)
			}
			if oldAppName != cfg.AppName {
				if err := webUtil.SetAppName(cfg.AppName); err != nil {
					fmt.Println(utils.ErrorWriter("Build failed."))
					return fmt.Errorf("failed to set web app name: %w", err)
				}
			}
		default:
			fmt.Println(utils.ErrorWriter("Build failed."))
			return fmt.Errorf("unsupported platform: %s", platform)
		}

		if err := flutter.Build(root, platform, clientName, cfg); err != nil {
			fmt.Println(utils.ErrorWriter("Build failed."))
			return fmt.Errorf("failed to build app: %w", err)
		}

		fmt.Println()
		fmt.Println(utils.SuccessWriter("Build finished successfully"))
		fmt.Println()

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
