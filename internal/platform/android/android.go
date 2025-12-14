package android

import (
	"errors"
	"flumint/internal/config"
	"flumint/internal/utils"
	"fmt"
	"os"
	"regexp"
)

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
	reg := regexp.MustCompile(`applicationId\s*"?"(.*?)"`)
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

func NewBundleStrategy(cfg config.AndroidConfig) (BundleIdStrategy, error) {
	hasGroovy := utils.FileExists(cfg.GradleGroovy)
	hasKts := utils.FileExists(cfg.GradleKts)

	if hasGroovy && hasKts {
		return nil, errors.New("both build.gradle and build.gradle.kts exist; remove one")
	}
	if hasGroovy {
		return GroovyStrategy{Path: cfg.GradleGroovy}, nil
	}
	if hasKts {
		return KotlinStrategy{Path: cfg.GradleKts}, nil
	}
	return nil, errors.New("no gradle build file found")
}

type Android struct {
	Config   config.AndroidConfig
	Strategy BundleIdStrategy
}

func NewAndroid(root string) *Android {
	return &Android{
		Config: config.DefaultAndroidConfig(root),
	}
}

func (a *Android) GetAppName() (string, error) {
	content, err := os.ReadFile(a.Config.ManifestMain)
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
		a.Config.ManifestMain,
		a.Config.ManifestDebug,
		a.Config.ManifestProfile,
	}

	for _, file := range files {
		utils.ReplaceInFileRegex(file,
			`android:label="(.*?)"`,
			fmt.Sprintf(`android:label="%s"`, name),
		)
	}

	return nil
}

func (a *Android) GetBundleId() (string, error) {
	return a.Strategy.ReadBundleId()
}

func (a *Android) SetBundleId(id string) error {
	if err := a.Strategy.WriteBundleId(id); err != nil {
		return err
	}
	return nil
}
