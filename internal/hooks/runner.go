package hooks

import (
	"fmt"
	"strings"

	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/git"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
)

// Run executes a list of commands for a specific hook sequentially (for now)
func Run(hookName string, cfg *config.Config) error {
	hookCfg, ok := cfg.Hooks[hookName]
	if !ok || len(hookCfg.Commands) == 0 {
		return nil // No hooks configured for this event
	}

	utils.Header(fmt.Sprintf("Running %s hooks...", hookName))

	var hasErrors bool

	if hookCfg.Parallel {
		hasErrors = runParallel(hookCfg.Commands)
	} else {
		hasErrors = runSequential(hookCfg.Commands)
	}

	fmt.Println()
	if hasErrors {
		utils.Error("%s hooks failed.", hookName)
		return fmt.Errorf("hook execution failed")
	}

	utils.Success("All hooks passed.")
	return nil
}

func runSequential(cmds []config.HookCmd) bool {
	hasErrors := false
	for _, cmd := range cmds {
		runStr, shouldSkip := processStagedFiles(cmd.Run)
		if shouldSkip {
			continue // Skip this command because it requires {staged_files} but none exist
		}

		displayName := cmd.Name
		if displayName == "" {
			displayName = runStr
		}

		utils.ProcessName(displayName)
		if cmd.Name != "" {
			utils.CommandLog(runStr)
		}
		fmt.Println()

		if err := utils.ExecuteShell(runStr, cmd.Env); err != nil {
			utils.Error("%s failed", displayName)
			hasErrors = true
			break // Stop on first error in sequential mode
		}

		utils.Success("Success")
	}
	return hasErrors
}

func runParallel(cmds []config.HookCmd) bool {
	utils.Info("Executing hooks in parallel...")

	type result struct {
		displayName string
		runStr      string
		output      string
		err         error
		skipped     bool
	}

	resCh := make(chan result, len(cmds))

	for _, cmd := range cmds {
		go func(c config.HookCmd) {
			runStr, shouldSkip := processStagedFiles(c.Run)
			if shouldSkip {
				resCh <- result{
					displayName: "Skipped (no staged files)",
					runStr:      c.Run,
					skipped:     true,
				}
				return
			}

			displayName := c.Name
			if displayName == "" {
				displayName = runStr
			}

			out, err := utils.ExecuteShellCaptured(runStr, c.Env)
			resCh <- result{
				displayName: displayName,
				runStr:      runStr,
				output:      out,
				err:         err,
			}
		}(cmd)
	}

	hasErrors := false
	for i := 0; i < len(cmds); i++ {
		res := <-resCh

		if res.skipped {
			continue
		}

		utils.ProcessName(res.displayName)
		if res.displayName != res.runStr {
			utils.CommandLog(res.runStr)
		}
		fmt.Println()

		// Print captured output if any
		if res.output != "" {
			fmt.Print(res.output)
		}

		if res.err != nil {
			utils.Error("%s failed: %v", res.displayName, res.err)
			hasErrors = true
		} else {
			utils.Success("Success")
		}
	}

	return hasErrors
}

// processStagedFiles replaces {staged_files} placeholder.
// Returns the processed run string, and a boolean true if the command should be skipped
// (which happens if {staged_files} is used but there are no staged files to pass).
func processStagedFiles(runStr string) (string, bool) {
	if !strings.Contains(runStr, "{staged_files}") {
		return runStr, false
	}

	stagedFiles, err := git.GetStagedFiles()
	// If git fails or there are no files, and the user explicitly requested {staged_files},
	// we skip the command entirely to save time.
	if err != nil || len(stagedFiles) == 0 {
		return runStr, true
	}

	// Replace {staged_files} with a space-separated string of the files
	filesStr := strings.Join(stagedFiles, " ")
	return strings.ReplaceAll(runStr, "{staged_files}", filesStr), false
}
