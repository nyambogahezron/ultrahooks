package cmd

import (
	"github.com/nyambogahezron/ultrahooks/internal/hooks"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall hooks",
	Long:  `Removes hooks created by UltraHooks from the .git/hooks directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := hooks.Uninstall(); err != nil {
			utils.Fatal("Failed to uninstall hooks: %v", err)
		}
		utils.Success("Uninstallation complete.")
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
