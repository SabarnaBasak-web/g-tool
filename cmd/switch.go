/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gtool/internal/git"
	"strings"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch [branch_name] and sync the branch",
	Short: "Switch to a existing branch and pull all the changes from the given branch",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		// Step 1: checkout
		_, errOut, err := git.RunGitCommandWithOutput("checkout", branchName)
		if err != nil {
			cmd.Println("❌ Failed to switch branch:", errOut)
			return
		}

		// Step 2: pull
		out, errOut, err := git.RunGitCommandWithOutput("pull", "origin", branchName)
		if err != nil {
			cmd.Println("❌ Pull failed:", errOut)
			return
		}

		// Step 3: Custom handling
		if strings.Contains(out, "Already up to date") {
			cmd.Println("✅ Already up to date")
			return
		}

		if strings.Contains(out, "Fast-forward") {
			cmd.Println("🔄 Branch updated (fast-forward)")
			return
		}

		cmd.Println("✅ Switched to:", branchName, "and synced")

	},
}

func init() {
	branchCmd.AddCommand(switchCmd)
}
