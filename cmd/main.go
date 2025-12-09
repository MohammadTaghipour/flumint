package main

import (
	"flumint/internal/android"
	"flumint/internal/config"
	"fmt"
)

func main() {
	config := config.Config{
		AndroidConfig: config.AndroidConfig{
			BuildGradlePath:     "android/app/build.gradle",
			ManifestPath:        "android/app/src/main/AndroidManifest.xml",
			ManifestDebugPath:   "android/app/src/debug/AndroidManifest.xml",
			ManifestProfilePath: "android/app/src/profile/AndroidManifest.xml",
			ActivityPath:        "android/app/src/main/",
		},
	}

	android := android.Android{
		Config: config.AndroidConfig,
	}

	if err := android.ChangePackageName("apps"); err != nil {
		fmt.Println("error on package rename")
	}
}
