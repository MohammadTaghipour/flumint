package service

import (
	"fmt"
	"time"

	"github.com/MohammadTaghipour/flumint/internal/assets"
	"github.com/MohammadTaghipour/flumint/internal/client"
	"github.com/MohammadTaghipour/flumint/internal/config"
	"github.com/MohammadTaghipour/flumint/internal/platform/android"
	"github.com/MohammadTaghipour/flumint/internal/platform/web"
	"github.com/MohammadTaghipour/flumint/internal/utils"
	"github.com/spf13/cobra"
)

func RunCheckout(cmd *cobra.Command) error {
	fmt.Println()

	s := newSpinner(" Running Flumint checkout...")
	s.Start()
	time.Sleep(time.Second)
	s.Stop()

	clientName, _ := cmd.Flags().GetString("client")
	root, _ := cmd.Flags().GetString("path")

	clientPath, err := client.Resolve(root, clientName)
	if err != nil || clientPath == "" {
		return checkoutFail("failed to resolve client", err)
	}
	fmt.Printf("client path resolved: %s\n", clientPath)

	cfg, err := config.Load(clientPath)
	if err != nil {
		return checkoutFail("failed to load config", err)
	}
	fmt.Println("config.json detected")
	fmt.Println(utils.InfoWriter(fmt.Sprintf("AppName     : %s", cfg.AppName)))
	fmt.Println(utils.InfoWriter(fmt.Sprintf("PackageName : %s", cfg.PackageName)))

	s = newSpinner(" Injecting client assets and files...")
	s.Start()
	time.Sleep(2 * time.Second)

	if err := assets.Inject(root, clientPath); err != nil {
		s.Stop()
		return checkoutFail("failed to inject assets", err)
	}
	s.Stop()
	fmt.Println("Assets injected successfully")

	if err := syncCheckoutAndroid(root, cfg); err != nil {
		return err
	}

	if err := syncCheckoutWeb(root, cfg); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println(utils.SuccessWriter("Checking out finished successfully"))
	fmt.Println()

	return nil
}

func syncCheckoutAndroid(root string, cfg *config.ClientConfig) error {

	androidUtil := android.NewAndroid(root)

	currAppName, err := androidUtil.GetAppName()
	if err != nil {
		return checkoutFail("failed to fetch android app name", err)
	}
	fmt.Printf("current app name is: %s\n", utils.InfoWriter(currAppName))

	if currAppName != cfg.AppName {
		if err := androidUtil.SetAppName(cfg.AppName); err != nil {
			return checkoutFail("failed to set android app name", err)
		}
		fmt.Printf("app name changed from %s to %s\n",
			utils.InfoWriter(currAppName),
			utils.InfoWriter(cfg.AppName))
	} else {
		fmt.Println("no changes in app name.")
	}

	currPackageName, err := androidUtil.GetPackageName()
	if err != nil {
		return checkoutFail("failed to fetch android package name", err)
	}
	fmt.Printf("current package name is: %s\n", utils.InfoWriter(currPackageName))

	if currPackageName != cfg.PackageName {
		if err := androidUtil.SetPackageName(cfg.PackageName); err != nil {
			return checkoutFail("failed to set android package name", err)
		}
		if err := androidUtil.SetPackageNameInActivities(cfg.PackageName); err != nil {
			return checkoutFail("failed to update package name in java/kotlin files", err)
		}
		if err := androidUtil.SetPackageNameInManifest(cfg.PackageName); err != nil {
			return checkoutFail("failed to update package name in manifest files", err)
		}
		fmt.Printf("package name updated to %s\n", utils.InfoWriter(cfg.PackageName))
	} else {
		fmt.Println("no changes in package name.")
	}

	return nil
}

func syncCheckoutWeb(root string, cfg *config.ClientConfig) error {

	webUtil := web.NewWeb(root)

	currAppName, err := webUtil.GetAppName()
	if err != nil {
		return checkoutFail("failed to fetch web app name", err)
	}
	fmt.Printf("current web app name is: %s\n", utils.InfoWriter(currAppName))

	if currAppName != cfg.AppName {
		if err := webUtil.SetAppName(cfg.AppName); err != nil {
			return checkoutFail("failed to set web app name", err)
		}
		fmt.Printf("web app name changed to %s\n", utils.InfoWriter(cfg.AppName))
	} else {
		fmt.Println("no changes in web app name.")
	}

	if err := webUtil.SetManifestInfo(cfg.AppName, cfg.AppDescription); err != nil {
		fmt.Println(utils.ErrorWriter(fmt.Sprintf("failed to set web info: %v", err)))
	} else {
		fmt.Println("web app info updated in Manifest.json")
	}

	return nil
}

func checkoutFail(message string, err error) error {
	fmt.Println(utils.ErrorWriter("Checkout failed."))
	if err != nil {
		return fmt.Errorf("%s: %w", message, err)
	}
	return fmt.Errorf("%s", message)
}
