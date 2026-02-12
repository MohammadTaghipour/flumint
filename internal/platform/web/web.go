package web

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Web struct {
	Config Config
}

func DefaultWebConfig(root string) Config {
	return Config{
		IndexHTMLPath: filepath.Join(root, "web", "index.html"),
	}
}

func NewWeb(root string) *Web {
	return &Web{
		Config: DefaultWebConfig(root),
	}
}

func (w *Web) GetAppName() (string, error) {
	data, err := os.ReadFile(w.Config.IndexHTMLPath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	titleRegex := regexp.MustCompile(`<title>(.*?)</title>`)

	for _, line := range lines {
		if titleRegex.MatchString(line) {
			match := titleRegex.FindStringSubmatch(line)
			if len(match) > 1 {
				return strings.TrimSpace(match[1]), nil
			}
		}
	}

	return "", errors.New("title tag not found in index.html")
}

func (w *Web) SetAppName(appName string) error {
	data, err := os.ReadFile(w.Config.IndexHTMLPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	titleRegex := regexp.MustCompile(`<title>(.*?)</title>`)
	updated := false

	for i, line := range lines {
		if titleRegex.MatchString(line) {
			lines[i] = fmt.Sprintf("  <title>%s</title>", appName)
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("title tag not found in index.html")
	}

	return os.WriteFile(w.Config.IndexHTMLPath, []byte(strings.Join(lines, "\n")), os.ModePerm)
}
