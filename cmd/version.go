package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "-v", "-V", "--version"},
	Short:   "Print the version number of UltraHooks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("UltraHooks %s %s/%s\n", Version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
