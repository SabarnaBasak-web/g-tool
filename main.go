package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runGitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: gtool <command>")
		return
	}

	switch args[1] {
	case "hello":
		fmt.Println("Hello from gtool 🚀")

	case "save":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gtool save <message>")
			return
		}
		msg := os.Args[2]
		if !IsGitInstalled() {
			fmt.Println("❌ Git is not installed. Please install Git first.")
			return
		}

		if !IsGitRepo() {
			fmt.Println("❌ Not a Git repository. Run `git init` first.")
			return
		}

		if !IsGitConfigured() {
			fmt.Println("❌ Git user not configured.")
			fmt.Println("Run:")
			fmt.Println(`  git config --global user.name "Your Name"`)
			fmt.Println(`  git config --global user.email "you@example.com"`)
			return
		}

		runGitCommand("add", ".")
		runGitCommand("commit", "-m", msg)
		runGitCommand("push")
	default:
		fmt.Println("Unknown command:", args[1])
	}
}
