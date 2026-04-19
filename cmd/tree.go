/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gtool/internal/git"

	"github.com/spf13/cobra"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show branch and commit tree",
	Long: `Shows a graphical view of your Git history across all branches.

It helps you quickly understand:
- where branches were created
- how they diverged
- where they were merged

Essentially a cleaner and more accessible version of 'git log --graph'.`,
	Run: func(cmd *cobra.Command, args []string) {
		git.RunGitCommand("log", "--oneline", "--graph", "--decorate", "--all")
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
