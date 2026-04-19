package cmd

import (
	"fmt"
	"gtool/internal/git"
	"os/exec"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save [message]",
	Short: "Add, commit, and push changes",
	Long: `stage all the existing files, commit with the message provided
push the changes to remote branch. For example:
gtool save "Message"`,
	Run: func(cmd *cobra.Command, args []string) {
		if !git.IsGitConfigSet() {
			return
		}

		if message == "" {
			fmt.Println("❌ Commit message required. Use -m")
			return
		}

		git.RunGitCommand("add", ".")
		git.RunGitCommand("commit", "-m", message)
		git.RunGitCommand("push")

		fmt.Println("✅ Changes saved and pushed")
	},
}

var message string

func runGit(cmd *cobra.Command, args ...string) {
	c := exec.Command("git", args...)
	c.Stdout = cmd.OutOrStdout()
	c.Stderr = cmd.ErrOrStderr()
	_ = c.Run()
}

func init() {
	saveCmd.Flags().StringVarP(
		&message,
		"message",        // long flag → --message
		"m",              // short flag → -m
		"",               // default value
		"Commit message", // description (shown in help)
	)
	rootCmd.AddCommand(saveCmd)
}
