/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gtool/internal/git"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [branch-name]",
	Short: "Create a new branch",
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		err := git.RunGitCommand("checkout", "-b", branchName)
		if err != nil {
			cmd.Println("❌ Failed to create branch:", err)
			return
		}

		cmd.Println("✅ Branch created:", branchName)

	},
}

func init() {
	branchCmd.AddCommand(createCmd)
}
