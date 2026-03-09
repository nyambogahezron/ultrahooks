package doctor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nyambogahezron/ultrahooks/internal/config"
	"github.com/nyambogahezron/ultrahooks/internal/git"
	"github.com/nyambogahezron/ultrahooks/internal/hooks"
	"github.com/nyambogahezron/ultrahooks/internal/utils"
)

func RunDiagnostics() {
	utils.Header("UltraHooks Diagnostics")

	issuesFound := 0

	// Check Git Repository
	gitDir, err := git.FindGitDir()
	if err != nil {
		utils.Error("Not inside a Git repository")
		issuesFound++
	} else {
		utils.Success("Git repository detected at %s", gitDir)
	}

	// Check Configuration Directory
	if _, err := os.Stat(config.ConfigDir); os.IsNotExist(err) {
		utils.Error("Configuration directory '%s' is missing. Run `ultrahooks init`.", config.ConfigDir)
		issuesFound++
	} else {
		utils.Success("Configuration directory '%s' exists", config.ConfigDir)
	}

	// Check config.yaml (Optional, but good to validate if exists)
	configPath := config.GetConfigPath()
	if _, err := os.Stat(configPath); err == nil {
		if _, err := config.Load(); err != nil {
			utils.Error("Failed to parse '%s': %v", configPath, err)
			issuesFound++
		} else {
			utils.Success("✔ Configuration file '%s' is valid", config.ConfigFile)
		}
	} else {
		utils.Info("Configuration file '%s' is not present (optional).", config.ConfigFile)
	}

	// Check Script Permissions
	if entries, err := os.ReadDir(config.ConfigDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			ext := filepath.Ext(entry.Name())
			if ext == ".sh" || ext == ".bat" || ext == ".ps1" {
				scriptPath := filepath.Join(config.ConfigDir, entry.Name())
				info, err := os.Stat(scriptPath)
				if err == nil {
					// Check if executable (specifically on Unix systems)
					if info.Mode()&0111 == 0 && ext == ".sh" {
						utils.Error(" Script '%s' is not executable. Run `chmod +x %s`", scriptPath, scriptPath)
						issuesFound++
					} else {
						utils.Success("✔ Script '%s' has correct permissions", scriptPath)
					}
				}
			}
		}
	}

	// Check Git Hook Proxy Linkage
	if gitDir != "" {
		cfg, _ := config.Load()
		if cfg != nil {
			for hookName := range cfg.Hooks {
				if hooks.IsValidHook(hookName) {
					hookPath := filepath.Join(gitDir, "hooks", hookName)
					if _, err := os.Stat(hookPath); os.IsNotExist(err) {
						utils.Error(" Git proxy script missing for active hook '%s'. Run `ultrahooks install`.", hookName)
						issuesFound++
					} else {
						utils.Success("✔ Git proxy script exists for '%s'", hookName)
					}
				}
			}
		}
	}

	fmt.Println()
	if issuesFound > 0 {
		utils.Error("Found %d issue(s) that may prevent UltraHooks from functioning correctly.", issuesFound)
	} else {
		utils.Success("No issues found! Your setup is healthy.")
	}
}
