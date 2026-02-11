package assets

import (
	"flumint/internal/utils"
	"fmt"
	"path/filepath"
)

// Inject copies client assets into lib/assets and android/google-services.json
func Inject(clientPath string) error {
	srcAssets := filepath.Join(clientPath, "assets")
	destAssets := "lib/assets"

	// copy assets
	if err := utils.CopyDirectory(srcAssets, destAssets); err != nil {
		return fmt.Errorf("failed to copy assets: %v", err)
	}

	// copy google-services.json for Android
	srcGoogle := filepath.Join(clientPath, "android", "google-services.json")
	destGoogle := filepath.Join("android", "app", "google-services.json")
	if err := utils.CopyFile(srcGoogle, destGoogle); err != nil {
		return fmt.Errorf("failed to copy google-services.json: %v", err)
	}

	return nil
}
