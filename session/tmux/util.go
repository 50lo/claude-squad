package tmux

import (
	"fmt"
	"os"
	"os/exec"

	"claude-squad/log"
)

// skipTmuxCheck can be set in tests to bypass the tmux presence check.
var skipTmuxCheck bool

// checkTmux checks if tmux is installed and reachable in PATH.
func checkTmux() error {
	if skipTmuxCheck {
		return nil
	}
	pathEnv := os.Getenv("PATH")
	tmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		if log.ErrorLog != nil {
			log.ErrorLog.Printf("tmux not found in PATH: %s", pathEnv)
		}
		return fmt.Errorf("tmux is not installed or not in PATH: %w", err)
	}
	if log.InfoLog != nil {
		log.InfoLog.Printf("tmux found at %s (PATH=%s)", tmuxPath, pathEnv)
	}
	return nil
}
