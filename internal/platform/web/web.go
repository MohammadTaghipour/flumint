package web

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MohammadTaghipour/flumint/internal/utils"
)

type Web struct {
	Config Config
}

func DefaultWebConfig(root string) Config {
	return Config{
		IndexHTMLPath: filepath.Join(root, "web", "index.html"),
		ManifestPath:  filepath.Join(root, "web", "manifest.json"),
	}
}

func NewWeb(root string) *Web {
	return &Web{
		Config: DefaultWebConfig(root),
	}
}

func (w *Web) GetAppName() (string, error) {
	path := w.Config.IndexHTMLPath
	if !utils.FileExists(path) {
		return "", fmt.Errorf("file not found %s", path)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read file %s. %w", path, err)
	}

	titleRegex := regexp.MustCompile(`<title>(.*?)</title>`)
	match := titleRegex.FindStringSubmatch(string(content))
	if len(match) > 1 {
		return strings.TrimSpace(match[1]), nil
	}

	return "", errors.New("title tag not found in index.html")
}

func (w *Web) SetAppName(appName string) error {
	path := w.Config.IndexHTMLPath
	if !utils.FileExists(path) {
		return fmt.Errorf("file not found %s", path)
	}
	if err := utils.ReplaceInFileRegex(
		path,
		`<title>(.*?)</title>`,
		fmt.Sprintf("<title>%s</title>", appName),
	); err != nil {
		return fmt.Errorf("cannot change web app name in %s. %w", path, err)
	}

	return nil
}

func (w *Web) SetManifestInfo(name, description string) error {
	path := w.Config.ManifestPath
	if !utils.FileExists(path) {
		return fmt.Errorf("file not found %s.", path)
	}

	if err := utils.ReplaceInFileRegex(
		path,
		`"name"\s*:\s*"[^"]*"`,
		fmt.Sprintf(`"name": "%s"`, name),
	); err != nil {
		return fmt.Errorf("cannot change web app name in %s. %w", path, err)

	}
	if err := utils.ReplaceInFileRegex(
		path,
		`"short_name"\s*:\s*"[^"]*"`,
		fmt.Sprintf(`"short_name": "%s"`, name),
	); err != nil {
		return fmt.Errorf("cannot change web app name in %s. %w", path, err)

	}

	if err := utils.ReplaceInFileRegex(
		path,
		`"description"\s*:\s*"[^"]*"`,
		fmt.Sprintf(`"description": "%s"`, description),
	); err != nil {
		return fmt.Errorf("cannot change web app description in %s. %w", path, err)
	}

	return nil
}
