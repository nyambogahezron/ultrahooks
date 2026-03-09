package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/hooks"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var removeAll bool

var removeCmd = &cobra.Command{
	Use:   "remove [hook-name]...",
	Short: "Remove git hooks",
	Long:  `Removes hooks from .git/hooks and deletes their custom script templates from .ultrahooks/`,
	ValidArgs: hooks.StandardHooks,
	Run: func(cmd *cobra.Command, args []string) {
		if removeAll {
			// Unwire all Git hooks
			if err := hooks.Uninstall(); err != nil {
				utils.Fatal("Failed to uninstall hooks: %v", err)
			}
			
			// Delete all script files in .ultrahooks
			removedScripts := 0
			for _, hookName := range hooks.StandardHooks {
				scriptPath := filepath.Join(config.ConfigDir, fmt.Sprintf("%s.sh", hookName))
				if err := os.Remove(scriptPath); err == nil {
					utils.Success("Deleted script: %s", scriptPath)
					removedScripts++
				}
			}
			if removedScripts == 0 {
				utils.Info("No custom shell scripts found to delete in .ultrahooks/")
			}
			return
		}

		if len(args) == 0 {
			utils.Fatal("You must specify hook names to remove, or use --all. e.g. ultrahooks remove pre-commit")
		}

		for _, hookName := range args {
			if !hooks.IsValidHook(hookName) {
				utils.Error("%s is not a recognized standard Git hook, skipping...", hookName)
				continue
			}

			// 1. Unwire Git proxy script
			removed, err := hooks.UninstallSingleHook(hookName)
			if err != nil {
				utils.Error("Failed to remove Git hook proxy for %s: %v", hookName, err)
			} else if !removed {
				utils.Info("No Git proxy script found for %s", hookName)
			}

			// 2. Delete the custom shell script
			scriptPath := filepath.Join(config.ConfigDir, fmt.Sprintf("%s.sh", hookName))
			if err := os.Remove(scriptPath); err == nil {
				utils.Success("Deleted script: %s", scriptPath)
			} else if os.IsNotExist(err) {
				utils.Info("No custom script found at %s", scriptPath)
			} else {
				utils.Error("Failed to delete %s: %v", scriptPath, err)
			}
		}
	},
}

func init() {
	removeCmd.Flags().BoolVarP(&removeAll, "all", "a", false, "Remove all UltraHooks proxy scripts and custom scripts")
	rootCmd.AddCommand(removeCmd)
}
