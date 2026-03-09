package hooks

import (
	"fmt"

	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
)

// Run executes a list of commands for a specific hook sequentially
func Run(hookName string, cfg *config.Config) error {
	cmds, ok := cfg.Hooks[hookName]
	if !ok || len(cmds) == 0 {
		return nil // No hooks configured for this event
	}

	utils.Header(fmt.Sprintf("Running %s hooks...", hookName))

	for _, cmd := range cmds {
		displayName := cmd.Name
		if displayName == "" {
			displayName = cmd.Run
		}

		utils.ProcessName(displayName)
		if cmd.Name != "" {
			utils.CommandLog(cmd.Run)
		}
		fmt.Println()

		if err := utils.ExecuteShell(cmd.Run); err != nil {
			utils.Error("%s failed", displayName)
			return err
		}

		utils.Success("Success")
	}

	fmt.Println()
	utils.Success("All hooks passed.")
	return nil
}
