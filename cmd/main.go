package main

import (
	"flumint/internal/android"
	"flumint/internal/config"
	"fmt"
)

func main() {
	workDir := "D:\\Programming\\FlutterProjects\\rubika_downloader"
	// workDir, err := os.Getwd()
	// if err != nil {
	// 	panic("error when getting current directory")
	// }

	config := config.Config{
		WorkingDir:    workDir,
		AndroidConfig: config.DefaultAndroidConfig(workDir),
	}

	strategy, err := android.NewBundleStrategy(config.AndroidConfig)
	if err != nil {
		fmt.Println("error on strategy", err)
	}
	android := android.Android{
		Config:   config.AndroidConfig,
		Strategy: strategy}

	if err := android.SetBundleId("com.example.app"); err != nil {
		fmt.Println("error on package rename", err)
	}
}
