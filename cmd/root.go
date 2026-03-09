package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"
)

var rootCmd = &cobra.Command{
	Use:     "ultrahooks",
	Short:   "A fast, lightweight, universal Git hooks manager",
	Long:    `UltraHooks is a language-agnostic Git hooks manager designed to be simple and fast.`,
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		vCap, _ := cmd.Flags().GetBool("V")

		if v || vCap {
			fmt.Printf("UltraHooks %s %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
	rootCmd.Flags().BoolP("V", "V", false, "Print version information")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.Error(err.Error())
		os.Exit(1)
	}
}
