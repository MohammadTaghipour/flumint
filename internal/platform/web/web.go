package web

import (
	"errors"
	"flumint/internal/config"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Web struct {
	Config config.WebConfig
}

func NewWeb(root string) *Web {
	return &Web{
		Config: config.DefaultWebConfig(root),
	}
}

func (w *Web) GetAppName() (string, error) {
	data, err := os.ReadFile(w.Config.IndexHTML)
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
	data, err := os.ReadFile(w.Config.IndexHTML)
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

	return os.WriteFile(w.Config.IndexHTML, []byte(strings.Join(lines, "\n")), os.ModePerm)
}

func (w *Web) GetBundleId() string {
	return "Web platform doesn't have bundleIdentifier."
}

func (w *Web) SetBundleId(bundleId string) string {
	return "Web platform doesn't have bundleIdentifier."
}
