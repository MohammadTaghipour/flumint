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

	data, err := os.ReadFile(w.Config.IndexHTMLPath)
	if err != nil {
		return fmt.Errorf("cannot read file %s. %w", path, err)
	}

	lines := strings.Split(string(data), "\n")
	titleRegex := regexp.MustCompile(`<title>(.*?)</title>`)
	updated := false

	for i, line := range lines {
		if titleRegex.MatchString(line) {
			lines[i] = fmt.Sprintf("\t<title>%s</title>", appName)
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("title tag not found in index.html")
	}

	if err := os.WriteFile(w.Config.IndexHTMLPath, []byte(strings.Join(lines, "\n")), os.ModePerm); err != nil {
		return fmt.Errorf("cannot write file %s. %w", path, err)
	}

	return nil
}
