package assets

import (
	"fmt"

	"github.com/MohammadTaghipour/flumint/internal/utils"
)

// Inject copies client files into project by keeping its file tree
func Inject(clientPath string) error {
	if !utils.DirectoryExists(clientPath) {
		return fmt.Errorf("client path does not exist: %s", clientPath)
	}

	return utils.CopyDirectory(clientPath, "./")

	//srcAssets := filepath.Join(clientPath, "assets")
	//destAssets := "assets"
	//
	//// copy assets
	//if err := utils.CopyDirectory(srcAssets, destAssets); err != nil {
	//	return fmt.Errorf("can not copy assets: %v", err)
	//}
	//
	//// copy google-services.json for Android
	//srcGoogle := filepath.Join(clientPath, "android", "google-services.json")
	//destGoogle := filepath.Join("android", "app", "google-services.json")
	//if err := utils.CopyFile(srcGoogle, destGoogle); err != nil {
	//	return fmt.Errorf("can not copy google-services.json: %v", err)
	//}
	//
	//return nil
}
