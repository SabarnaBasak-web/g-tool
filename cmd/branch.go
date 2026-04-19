/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches",
	Long:  `Create, rename, and manage Git branches easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("branch called")
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
