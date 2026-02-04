package flutter

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func (f *Flutter) GetVersion() (*VersionInfo, error) {
	cmd := exec.Command("flutter", "--version")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error in getting flutter version: %v", err)
	}

	return parse(string(out))
}

type VersionInfo struct {
	Version  string `json:"version"`
	Channel  string `json:"channel"`
	Dart     string `json:"dart"`
	DevTools string `json:"dev_tools"`
}

// parses flutter --version output and returns FlutterVersion
func parse(output string) (*VersionInfo, error) {
	if strings.TrimSpace(output) == "" {
		return nil, errors.New("empty flutter version output")
	}

	v := &VersionInfo{}
	var errs []string

	// version
	versionMatch := regexp.MustCompile(`Flutter\s+([0-9]+\.[0-9]+\.[0-9]+)`).
		FindStringSubmatch(output)
	if len(versionMatch) > 1 {
		v.Version = versionMatch[1]
	} else {
		errs = append(errs, "version not found")
	}

	// channel
	channelMatch := regexp.MustCompile(`channel\s+(\w+)`).FindStringSubmatch(output)
	if len(channelMatch) > 1 {
		v.Channel = channelMatch[1]
	} else {
		errs = append(errs, "channel not found")
	}

	// Dart version
	dartMatch := regexp.MustCompile(`Dart\s+([0-9]+\.[0-9]+\.[0-9]+)`).
		FindStringSubmatch(output)
	if len(dartMatch) > 1 {
		v.Dart = dartMatch[1]
	} else {
		errs = append(errs, "dart version not found")
	}

	// DevTools version (optional - no error if not found)
	devToolsMatch := regexp.MustCompile(`DevTools\s+([0-9]+\.[0-9]+\.[0-9]+)`).
		FindStringSubmatch(output)
	if len(devToolsMatch) > 1 {
		v.DevTools = devToolsMatch[1]
	}

	if len(errs) > 0 {
		return v, errors.New(strings.Join(errs, "; "))
	}

	return v, nil
}
