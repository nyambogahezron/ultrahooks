package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hooks system",
	Long:  `Creates the .ultrahooks directory with config.yaml and template hooks.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Header("Welcome to UltraHooks!")
		fmt.Println("Let's set up your environment.")

		var languages []string
		prompt := &survey.MultiSelect{
			Message: "What languages does your project use?",
			Options: []string{"Go", "Node.js", "Python", "Rust", "Other"},
		}

		err := survey.AskOne(prompt, &languages)
		if err != nil {
			utils.Fatal("Initialization cancelled.")
		}

		if err := config.CreateDefault(languages); err != nil {
			utils.Fatal("Failed to create config: %v", err)
		}

		utils.Success("Created %s with tailored hooks", config.ConfigFile)
		utils.Info("Awesome! Now run `ultrahooks install` to wire them.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
