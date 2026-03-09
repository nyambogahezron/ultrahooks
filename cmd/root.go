package cmd

import (
	"os"

	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ultrahooks",
	Short: "A fast, lightweight, universal Git hooks manager",
	Long:  `UltraHooks is a language-agnostic Git hooks manager designed to be simple and fast.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.Error(err.Error())
		os.Exit(1)
	}
}
