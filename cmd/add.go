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

var addCmd = &cobra.Command{
	Use:       "add [hook-name]...",
	Short:     "Add new git hooks",
	Long:      `Scaffolds new executable templates under .ultrahooks/ and instantly wires them to .git/hooks.`,
	ValidArgs: hooks.StandardHooks,
	Args:      cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure .ultrahooks exists regardless of the loop
		if err := os.MkdirAll(config.ConfigDir, 0755); err != nil {
			utils.Fatal("Failed to create %s directory: %v", config.ConfigDir, err)
		}

		for _, hookName := range args {
			if !hooks.IsValidHook(hookName) {
				utils.Error("%s is not a recognized standard Git hook, skipping...", hookName)
				continue
			}

			scriptPath := filepath.Join(config.ConfigDir, fmt.Sprintf("%s.sh", hookName))

			// Create the bash script if it doesn't already exist
			if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
				content := fmt.Sprintf("#!/bin/sh\n# Commands executed on %s\necho \"%s triggered\"\n", hookName, hookName)
				if err := os.WriteFile(scriptPath, []byte(content), 0755); err != nil {
					utils.Error("Failed to create script %s: %v", scriptPath, err)
					continue
				}
				utils.Success("Scaffolded %s", scriptPath)
			} else {
				utils.Warning("Script %s already exists, skipping scaffold.", scriptPath)
			}

			// Wire it up in Git
			if err := hooks.InstallSingleHook(hookName); err != nil {
				utils.Error("Failed to install Git hook proxy for %s: %v", hookName, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
