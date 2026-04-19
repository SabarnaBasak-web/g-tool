package git

import (
	"fmt"
	"os/exec"
)

func IsGitInstalled() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	return cmd.Run() == nil
}

func IsGitConfigured() bool {
	nameCmd := exec.Command("git", "config", "--get", "user.name")
	emailCmd := exec.Command("git", "config", "--get", "user.email")

	if err := nameCmd.Run(); err != nil {
		return false
	}
	if err := emailCmd.Run(); err != nil {
		return false
	}
	return true
}

func IsGitConfigSet() bool {
	if !IsGitInstalled() {
		fmt.Println("❌ Git is not installed. Please install Git first.")
		return false
	}

	if !IsGitRepo() {
		fmt.Println("❌ Not a Git repository. Run `git init` first.")
		return false
	}

	if !IsGitConfigured() {
		fmt.Println("❌ Git user not configured.")
		fmt.Println("Run:")
		fmt.Println(`  git config --global user.name "Your Name"`)
		fmt.Println(`  git config --global user.email "you@example.com"`)
		return false
	}
	return true
}
