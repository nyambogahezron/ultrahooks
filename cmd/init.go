package cmd

import (
	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hooks system",
	Long:  `Creates the .ultrahooks directory with config.yaml and template hooks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.CreateDefault(); err != nil {
			utils.Fatal("Failed to create config: %v", err)
		}
		utils.Success("Created %s with default hooks", config.ConfigFile)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
