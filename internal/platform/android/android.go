package android

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MohammadTaghipour/flumint/internal/utils"
)

type Android struct {
	config Config
}

func NewAndroid() *Android {
	root := "./"
	return &Android{
		config: Config{
			ProjectRootPath:     root,
			ManifestMainPath:    filepath.Join(append([]string{root}, strings.Split("android/app/src/main/AndroidManifest.xml", "/")...)...),
			ManifestDebugPath:   filepath.Join(append([]string{root}, strings.Split("android/app/src/debug/AndroidManifest.xml", "/")...)...),
			ManifestProfilePath: filepath.Join(append([]string{root}, strings.Split("android/app/src/profile/AndroidManifest.xml", "/")...)...),
			GradleGroovyPath:    filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle", "/")...)...),
			GradleKtsPath:       filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle.kts", "/")...)...),
			ActivityPath:        filepath.Join(append([]string{root}, strings.Split("android/app/src/main", "/")...)...),
		},
	}
}

func (a *Android) getGradlePath() (string, error) {
	var gradlePath string
	if utils.FileExists(a.config.GradleGroovyPath) && utils.FileExists(a.config.GradleKtsPath) {
		return "", fmt.Errorf("both build.gradle and build.gradle.kts exist; remove one")
	}

	if utils.FileExists(a.config.GradleGroovyPath) {
		gradlePath = a.config.GradleGroovyPath
	} else if utils.FileExists(a.config.GradleKtsPath) {
		gradlePath = a.config.GradleKtsPath
	} else {
		return "", fmt.Errorf("build.gradle or build.gradle.kts not found")
	}

	return gradlePath, nil
}

func (a *Android) GetPackageName() (string, error) {
	gradlePath, err := a.getGradlePath()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(gradlePath)
	if err != nil {
		return "", fmt.Errorf("can not read file %s. %w", gradlePath, err)
	}

	reg := regexp.MustCompile(`applicationId\s*=?\s*"(.*)"`)
	match := reg.FindStringSubmatch(string(content))
	if match == nil {
		return "", fmt.Errorf("applicationId not found in %s", gradlePath)
	}
	return match[1], nil
}

func (a *Android) SetPackageName(newPackageName string) error {
	gradlePath, err := a.getGradlePath()
	if err != nil {
		return err
	}
	if err := utils.ReplaceInFileRegex(gradlePath, `applicationId\s*=?\s*"(.*)"`, newPackageName); err != nil {
		return fmt.Errorf("can not set package name in %s to %s", gradlePath, newPackageName)
	}
	return nil
}

func (a *Android) SetPackageNameInManifest(newPackageName string) error {
	manifests := []string{
		a.config.ManifestMainPath,
		a.config.ManifestDebugPath,
		a.config.ManifestProfilePath,
	}
	for _, path := range manifests {
		if utils.FileExists(path) {
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("can not read file %s. %w", path, err)
			}

			reg := regexp.MustCompile(`applicationId\s*=?\s*"(.*)"`)
			match := reg.FindStringSubmatch(string(content))
			if match != nil {
				replacement := fmt.Sprintf("package=\"%s\">", newPackageName)
				utils.ReplaceInFileRegex(a.config.ManifestMainPath, `(package=.*)`, replacement)
			}
		}
	}
	return nil
}
