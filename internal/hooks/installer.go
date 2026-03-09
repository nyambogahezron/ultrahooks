package hooks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/git"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
)

var StandardHooks = []string{
	"applypatch-msg", "pre-applypatch", "post-applypatch",
	"pre-commit", "pre-merge-commit", "prepare-commit-msg",
	"commit-msg", "post-commit", "pre-rebase", "post-checkout",
	"post-merge", "pre-push", "pre-receive", "update",
	"proc-receive", "post-receive", "post-update",
	"reference-transaction", "push-to-checkout", "pre-auto-gc",
	"post-rewrite", "sendemail-validate", "fsmonitor-watchman",
	"p4-changelist", "p4-prepare-changelist", "p4-post-changelist",
	"p4-pre-submit", "post-index-change",
}

// IsValidHook checks if a hookName is a standard Git hook
func IsValidHook(hookName string) bool {
	for _, h := range StandardHooks {
		if h == hookName {
			return true
		}
	}
	return false
}

// InstallSingleHook instantly wires up a single shell script for a Git hook
func InstallSingleHook(hookName string) error {
	gitDir, err := git.FindGitDir()
	if err != nil {
		return err
	}

	hooksDir := filepath.Join(gitDir, "hooks")
	if err := os.MkdirAll(hooksDir, 0755); err != nil {
		return err
	}

	hookPath := filepath.Join(hooksDir, hookName)
	script := fmt.Sprintf(HookTemplate, hookName)

	if err := os.WriteFile(hookPath, []byte(script), 0755); err != nil {
		return err
	}
	utils.Success("Installed script for %s", hookName)
	return nil
}

// Install hooks defined in configuration
func Install(cfg *config.Config) error {
	for hookName := range cfg.Hooks {
		if err := InstallSingleHook(hookName); err != nil {
			return err
		}
	}
	return nil
}

// UninstallSingleHook unwires a single Git hook if it was created by UltraHooks
func UninstallSingleHook(hookName string) (bool, error) {
	gitDir, err := git.FindGitDir()
	if err != nil {
		return false, err
	}

	hookPath := filepath.Join(gitDir, "hooks", hookName)
	data, err := os.ReadFile(hookPath)
	if err != nil {
		return false, nil // Doesn't exist or can't read
	}

	if strings.Contains(string(data), "ultrahooks run") {
		if err := os.Remove(hookPath); err == nil {
			utils.Success("Removed Git proxy script for: %s", hookName)
			return true, nil
		}
		return false, err
	}

	return false, nil
}

// Uninstall removed hooks
func Uninstall() error {
	removed := 0
	for _, hookName := range StandardHooks {
		ok, err := UninstallSingleHook(hookName)
		if err != nil {
			return err
		}
		if ok {
			removed++
		}
	}

	if removed == 0 {
		utils.Info("No UltraHooks scripts found to uninstall.")
	}

	return nil
}
