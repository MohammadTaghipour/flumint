package flutter

import (
	"fmt"
	"os/exec"
)

type Client interface {
	GetVersion() (*VersionInfo, error)
	RunDoctor() (string, error)
}

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) RunDoctor() (string, error) {
	cmd := exec.Command("flutter", "doctor")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("flutter doctor failed: %w", err)
	}
	return string(out), nil
}

func (c *CLI) GetVersion() (*VersionInfo, error) {
	cmd := exec.Command("flutter", "--version")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get flutter version: %w", err)
	}

	return parseVersion(string(out))
}
