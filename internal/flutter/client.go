package flutter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MohammadTaghipour/flumint/internal/config"
	"github.com/MohammadTaghipour/flumint/internal/utils"
	"gopkg.in/yaml.v3"
)

func IsFlutterProject(root string) (bool, error) {
	pubspecPath := filepath.Join(root, "pubspec.yaml")

	if !utils.FileExists(pubspecPath) {
		return false, fmt.Errorf("cannot access %s", pubspecPath)
	}

	data, err := os.ReadFile(pubspecPath)
	if err != nil {
		return false, fmt.Errorf("cannot read %s: %w", pubspecPath, err)
	}

	var p Pubspec
	if err := yaml.Unmarshal(data, &p); err != nil {
		return false, fmt.Errorf("invalid pubspec.yaml: %w", err)
	}

	if dep, ok := p.Dependencies["flutter"]; ok {
		if m, ok := dep.(map[string]interface{}); ok {
			if sdk, ok := m["sdk"]; ok && sdk == "flutter" {
				return true, nil
			}
		}
	}

	return false, nil
}

func RunDoctor() (string, error) {
	cmd := exec.Command("flutter", "doctor")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("flutter doctor error: %w", err)
	}
	return string(out), nil
}

func GetVersion() (*VersionInfo, error) {
	cmd := exec.Command("flutter", "--version")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("can not get flutter version: %w", err)
	}

	return parseVersion(string(out))
}

func Build(root, platform, clientName string, cfg *config.ClientConfig) error {
	var cmd *exec.Cmd

	dartDefines := []string{}
	for k, v := range cfg.DartDefines {
		dartDefines = append(dartDefines, fmt.Sprintf("--dart-define=%s=%s", k, v))
	}
	if clientName != "" {
		dartDefines = append(dartDefines, fmt.Sprintf("--dart-define=CLIENT=%s", clientName))
	}

	switch platform {
	case "android":
		args := append([]string{"build", "apk", "--release"}, dartDefines...)
		cmd = exec.Command("flutter", args...)
	case "web":
		args := append([]string{"build", "web", "--release"}, dartDefines...)
		cmd = exec.Command("flutter", args...)
	default:
		return fmt.Errorf("unsupported platform: %s", platform)
	}

	cmd.Dir = root

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running: %s %s\n", cmd.Path, strings.Join(cmd.Args[1:], " "))
	return cmd.Run()
}
