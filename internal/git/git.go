package git

import (
	"os"
	"os/exec"
)

func RunGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
