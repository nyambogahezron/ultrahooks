package cmd

import (
	"github.com/nyambogahezron/ultrahooks/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check the health of the UltraHooks setup",
	Long:  `Scans the project for .ultrahooks, config.yaml syntaxes, script permissions, and Git proxy links to diagnose issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		doctor.RunDiagnostics()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
