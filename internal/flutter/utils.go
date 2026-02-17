package flutter

import (
	"errors"
	"regexp"
	"strings"
)

var (
	versionRegex  = regexp.MustCompile(`Flutter\s+([0-9]+\.[0-9]+\.[0-9]+)`)
	channelRegex  = regexp.MustCompile(`channel\s+(\w+)`)
	dartRegex     = regexp.MustCompile(`Dart\s+([0-9]+\.[0-9]+\.[0-9]+)`)
	devToolsRegex = regexp.MustCompile(`DevTools\s+([0-9]+\.[0-9]+\.[0-9]+)`)
)

func parseVersion(output string) (*VersionInfo, error) {
	if strings.TrimSpace(output) == "" {
		return nil, errors.New("empty flutter version output")
	}

	v := &VersionInfo{}
	var errs []string

	if match := versionRegex.FindStringSubmatch(output); len(match) > 1 {
		v.Version = match[1]
	} else {
		errs = append(errs, "version not found")
	}

	if match := channelRegex.FindStringSubmatch(output); len(match) > 1 {
		v.Channel = match[1]
	} else {
		errs = append(errs, "channel not found")
	}

	if match := dartRegex.FindStringSubmatch(output); len(match) > 1 {
		v.Dart = match[1]
	} else {
		errs = append(errs, "dart version not found")
	}

	if match := devToolsRegex.FindStringSubmatch(output); len(match) > 1 {
		v.DevTools = match[1]
	}

	if len(errs) > 0 {
		return v, errors.New(strings.Join(errs, "; "))
	}

	return v, nil
}
