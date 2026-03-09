package cmd

import (
	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/hooks"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [hook]",
	Short: "Run hooks",
	Long:  `Executes the specified hook defined in .ultrahooks.yml sequentially.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hookName := args[0]

		cfg, err := config.Load()
		if err != nil {
			utils.Fatal("Failed to load config: %v", err)
		}

		if err := hooks.Run(hookName, cfg); err != nil {
			utils.Fatal("Hook execution failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
