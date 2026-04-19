package git

import (
	"os/exec"
)

func RunGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	return cmd.Run()
}
