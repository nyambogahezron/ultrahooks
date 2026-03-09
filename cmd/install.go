package cmd

import (
	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/hooks"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install hooks into git",
	Long:  `Detects the .git directory and creates hook scripts inside .git/hooks/.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			utils.Fatal("Failed to load config (run 'ultrahooks init' first): %v", err)
		}

		if err := hooks.Install(cfg); err != nil {
			utils.Fatal("Failed to install hooks: %v", err)
		}

		utils.Success("All hooks installed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
