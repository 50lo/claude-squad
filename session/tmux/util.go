package tmux

import (
	"fmt"
	"os/exec"
)

// skipTmuxCheck can be set in tests to bypass the tmux presence check.
var skipTmuxCheck bool

// checkTmux checks if tmux is installed and reachable in PATH.
func checkTmux() error {
	if skipTmuxCheck {
		return nil
	}
	if _, err := exec.LookPath("tmux"); err != nil {
		return fmt.Errorf("tmux is not installed or not in PATH: %w", err)
	}
	return nil
}
