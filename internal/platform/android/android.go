package android

import (
	"errors"
	"flumint/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func DefaultAndroidConfig(root string) Config {
	return Config{
		ProjectRootPath:      root,
		ManifestMainPath:     filepath.Join(append([]string{root}, strings.Split("android/app/src/main/AndroidManifest.xml", "/")...)...),
		ManifestDebugPath:    filepath.Join(append([]string{root}, strings.Split("android/app/src/debug/AndroidManifest.xml", "/")...)...),
		ManifestProfilePath:  filepath.Join(append([]string{root}, strings.Split("android/app/src/profile/AndroidManifest.xml", "/")...)...),
		GradleGroovyPath:     filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle", "/")...)...),
		GradleKtsPath:        filepath.Join(append([]string{root}, strings.Split("android/app/build.gradle.kts", "/")...)...),
		MainActivityRootPath: filepath.Join(append([]string{root}, strings.Split("android/app/src/main", "/")...)...),
	}
}

type BundleIdStrategy interface {
	ReadBundleId() (string, error)
	WriteBundleId(newId string) error
}

type GroovyStrategy struct {
	Path string
}

func (g GroovyStrategy) ReadBundleId() (string, error) {
	content, err := os.ReadFile(g.Path)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`applicationId\s*=?\s*"(.*?)"`)
	match := reg.FindStringSubmatch(string(content))
	if len(match) < 2 {
		return "", errors.New("applicationId not found in Groovy gradle")
	}
	return match[1], nil
}

func (g GroovyStrategy) WriteBundleId(newId string) error {
	return utils.ReplaceInFileRegex(
		g.Path,
		`applicationId\s*"(.*?)"`,
		fmt.Sprintf(`applicationId "%s"`, newId),
	)
}

type KotlinStrategy struct {
	Path string
}

func (k KotlinStrategy) ReadBundleId() (string, error) {
	content, err := os.ReadFile(k.Path)
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`applicationId\s*=\s*"(.*?)"`)
	match := reg.FindStringSubmatch(string(content))
	if len(match) < 2 {
		return "", errors.New("applicationId not found in Kotlin DSL gradle")
	}
	return match[1], nil
}

func (k KotlinStrategy) WriteBundleId(newId string) error {
	return utils.ReplaceInFileRegex(
		k.Path,
		`applicationId\s*=\s*"(.*?)"`,
		fmt.Sprintf(`applicationId = "%s"`, newId),
	)
}

func NewBundleStrategy(cfg *Config) (BundleIdStrategy, error) {
	hasGroovy := utils.FileExists(cfg.GradleGroovyPath)
	hasKts := utils.FileExists(cfg.GradleKtsPath)

	if hasGroovy && hasKts {
		return nil, errors.New("both build.gradle and build.gradle.kts exist; remove one")
	}
	if hasGroovy {
		return GroovyStrategy{Path: cfg.GradleGroovyPath}, nil
	}
	if hasKts {
		return KotlinStrategy{Path: cfg.GradleKtsPath}, nil
	}
	return nil, errors.New("no gradle build file found")
}

type Android struct {
	Config   Config
	Strategy BundleIdStrategy
}

func NewAndroid(root string) (*Android, error) {
	cfg := DefaultAndroidConfig(root)

	strategy, err := NewBundleStrategy(&cfg)
	if err != nil {
		return nil, err
	}

	return &Android{
		Config:   cfg,
		Strategy: strategy,
	}, nil
}

func (a *Android) GetAppName() (string, error) {
	content, err := os.ReadFile(a.Config.ManifestMainPath)
	if err != nil {
		return "", err
	}

	reg := regexp.MustCompile(`android:label="(.*?)"`)
	match := reg.FindStringSubmatch(string(content))
	if len(match) < 2 {
		return "", errors.New("app name not found")
	}
	return match[1], nil
}

func (a *Android) SetAppName(name string) error {
	files := []string{
		a.Config.ManifestMainPath,
		a.Config.ManifestDebugPath,
		a.Config.ManifestProfilePath,
	}

	for _, file := range files {
		if utils.FileExists(file) {
			if err := utils.ReplaceInFileRegex(
				file,
				`android:label="(.*?)"`,
				fmt.Sprintf(`android:label="%s"`, name),
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *Android) GetBundleId() (string, error) {
	if a.Strategy == nil {
		return "", errors.New("bundle strategy not initialized")
	}
	return a.Strategy.ReadBundleId()
}

func (a *Android) SetBundleId(id string) error {
	if a.Strategy == nil {
		return errors.New("bundle strategy not initialized")
	}
	return a.Strategy.WriteBundleId(id)
}
