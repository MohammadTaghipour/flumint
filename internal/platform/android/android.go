package android

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MohammadTaghipour/flumint/internal/utils"
)

type Android struct {
	config Config
}

func NewAndroid(root string) *Android {
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
	if len(match) < 2 {
		return "", fmt.Errorf("applicationId not found in %s", gradlePath)
	}
	return match[1], nil
}

func (a *Android) SetPackageName(newPackageName string) error {
	gradlePath, err := a.getGradlePath()
	if err != nil {
		return err
	}
	replacement := fmt.Sprintf(`applicationId = "%s"`, newPackageName)
	if err := utils.ReplaceInFileRegex(gradlePath, `applicationId\s*=?\s*"(.*)"`, replacement); err != nil {
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
		if !utils.FileExists(path) {
			continue
		}

		if err := utils.ReplaceInFileRegex(path, `package="[^"]*"`, fmt.Sprintf(`package="%s"`, newPackageName)); err != nil {
			return fmt.Errorf("cannot change manifest package %s: %w", path, err)
		}
	}

	return nil
}

func (a *Android) GetAppName() (string, error) {
	if !utils.FileExists(a.config.ManifestMainPath) {
		return "", fmt.Errorf("file not found %s", a.config.ManifestMainPath)
	}

	content, err := os.ReadFile(a.config.ManifestMainPath)
	if err != nil {
		return "", fmt.Errorf("can not read file %s. %w", a.config.ManifestMainPath, err)
	}

	reg := regexp.MustCompile(`android:label="(.*?)"`)
	match := reg.FindStringSubmatch(string(content))
	if len(match) < 2 {
		return "", fmt.Errorf("app name not found in %s", a.config.ManifestMainPath)
	}
	return match[1], nil

}

func (a *Android) SetAppName(name string) error {
	files := []string{
		a.config.ManifestMainPath,
		a.config.ManifestDebugPath,
		a.config.ManifestProfilePath,
	}

	for _, file := range files {
		if utils.FileExists(file) {
			if err := utils.ReplaceInFileRegex(
				file,
				`android:label="(.*?)"`,
				fmt.Sprintf(`android:label="%s"`, name),
			); err != nil {
				return fmt.Errorf("cannot change app name in %s. %w", file, err)
			}
		}
	}

	return nil
}

func (a *Android) SetPackageNameInActivities(newPackageName string) error {
	var javaKotlinFiles []string

	err := filepath.Walk(a.config.ActivityPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext == ".java" || ext == ".kt" {
			javaKotlinFiles = append(javaKotlinFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(javaKotlinFiles) == 0 {
		return nil
	}

	oldPackageName, err := a.extractPackageName(javaKotlinFiles[0])
	if err != nil {
		return err
	}

	if oldPackageName == newPackageName {
		return nil
	}

	type fileMove struct {
		oldPath string
		newPath string
	}

	var moves []fileMove

	for _, oldPath := range javaKotlinFiles {
		content, err := os.ReadFile(oldPath)
		if err != nil {
			return err
		}

		ext := filepath.Ext(oldPath)
		updated := updatePackageAndImports(string(content), oldPackageName, newPackageName, ext)

		baseType := "java"
		if ext == ".kt" {
			baseType = "kotlin"
		}

		baseSourcePath := filepath.Join(a.config.ActivityPath, baseType)
		newPackagePath := filepath.Join(
			baseSourcePath,
			filepath.FromSlash(strings.ReplaceAll(newPackageName, ".", "/")),
		)

		if err := os.MkdirAll(newPackagePath, os.ModePerm); err != nil {
			return err
		}

		fileName := filepath.Base(oldPath)
		newPath := filepath.Join(newPackagePath, fileName)

		if err := os.WriteFile(newPath, []byte(updated), 0644); err != nil {
			return err
		}

		moves = append(moves, fileMove{
			oldPath: oldPath,
			newPath: newPath,
		})
	}

	for _, m := range moves {
		if err := os.Remove(m.oldPath); err != nil {
			return err
		}
	}

	return utils.DeleteEmptyDirs(a.config.ActivityPath)
}

func (a *Android) extractPackageName(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "package ") {
			packagePart := strings.TrimPrefix(trimmedLine, "package ")
			packageName := strings.TrimSuffix(packagePart, ";")
			return strings.TrimSpace(packageName), nil
		}
	}

	return "", fmt.Errorf("package name not found in file: %s", filePath)
}

func updatePackageAndImports(
	content string,
	oldPkg string,
	newPkg string,
	ext string,
) string {

	lines := strings.Split(content, "\n")
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "package ") {
			if ext == ".kt" {
				result = append(result, "package "+newPkg)
			} else {
				result = append(result, "package "+newPkg+";")
			}
			continue
		}

		if strings.HasPrefix(trimmed, "import ") {
			importPath := strings.TrimPrefix(trimmed, "import ")
			importPath = strings.TrimSuffix(importPath, ";")
			importPath = strings.TrimSpace(importPath)

			if strings.HasPrefix(importPath, oldPkg+".") || importPath == oldPkg {
				newImport := strings.Replace(importPath, oldPkg, newPkg, 1)

				if ext == ".kt" {
					result = append(result, "import "+newImport)
				} else {
					result = append(result, "import "+newImport+";")
				}
				continue
			}
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}
