package config

import (
	"path/filepath"
	"strings"
)

type WebConfig struct {
	IndexHTML string
}

type AndroidConfig struct {
	ProjectRoot      string
	ManifestMain     string
	ManifestDebug    string
	ManifestProfile  string
	GradleGroovy     string
	GradleKts        string
	MainActivityRoot string // android/app/src/main/(java|kotlin)
}

func DefaultWebConfig(root string) WebConfig {
	return WebConfig{
		IndexHTML: filepath.Join(root, "web", "index.html"),
	}
}

func DefaultAndroidConfig(root string) AndroidConfig {
	return AndroidConfig{
		ProjectRoot:      root,
		ManifestMain:     filepath.Join(append([]string{root}, strings.Split("android/app/src/main/AndroidManifest.xml", "/")...)...),
		ManifestDebug:    filepath.Join(append([]string{root}, strings.Split("android/app/src/debug/AndroidManifest.xml", "/")...)...),
		ManifestProfile:  filepath.Join(append([]string{root}, strings.Split("android/app/src/profile/AndroidManifest.xml", "/")...)...),
		GradleGroovy:     filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle", "/")...)...),
		GradleKts:        filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle.kts", "/")...)...),
		MainActivityRoot: filepath.Join(append([]string{root}, strings.Split("android/app/src/main", "/")...)...),
	}
}

type Config struct {
	WorkingDir    string
	AndroidConfig AndroidConfig
}
