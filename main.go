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

		runGitCommand("add", ".")
		runGitCommand("commit", "-m", msg)
		runGitCommand("push")
	default:
		fmt.Println("Unknown command:", args[1])
	}
}
