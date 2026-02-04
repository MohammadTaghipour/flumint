package flutter

import (
	"fmt"
	"os/exec"
)

func (f *Flutter) RunDoctor() (string, error) {
	cmd := exec.Command("flutter", "doctor")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("flutter doctor failed: %v", err)
	}
	return string(out), nil
}
