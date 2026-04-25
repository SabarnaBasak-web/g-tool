package cmd

import (
	"gtool/internal/git"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename [old-name] [new-name]",
	Short: "Rename a branch",
	Args:  cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		oldName := args[0]
		newName := args[1]
		err := git.RunGitCommand("branch", "-m", oldName, newName)
		if err != nil {
			cmd.Println("❌ Failed to rename branch:", err)
			return
		}
		cmd.Println("✅ Branch renamed:", oldName, "→", newName)
	},
}

func init() {
	branchCmd.AddCommand(renameCmd)
}
