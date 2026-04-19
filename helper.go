package main

import (
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
