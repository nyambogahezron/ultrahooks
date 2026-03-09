package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	ConfigDir  = ".ultrahooks"
	ConfigFile = "config.yaml"
)

// GetConfigPath returns the full path to the configuration file
func GetConfigPath() string {
	return filepath.Join(ConfigDir, ConfigFile)
}

// Load reads and parses the configuration file
func Load() (*Config, error) {
	cfg := &Config{
		Hooks: make(map[string]HookConfig),
	}

	configPath := GetConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if err == nil {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, err
		}
	}

	if cfg.Hooks == nil {
		cfg.Hooks = make(map[string]HookConfig)
	}

	// Auto-discover shell scripts in the config directory to make config.yaml optional
	if entries, err := os.ReadDir(ConfigDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			ext := filepath.Ext(entry.Name())
			if ext == ".sh" || ext == ".bat" || ext == ".ps1" {
				hookName := entry.Name()[:len(entry.Name())-len(ext)] // Remove extension
				scriptPath := filepath.Join(ConfigDir, entry.Name())
				runStr := "./" + scriptPath

				hookCfg := cfg.Hooks[hookName]
				found := false
				for _, cmd := range hookCfg.Commands {
					if cmd.Run == runStr {
						found = true
						break
					}
				}

				if !found {
					// Prepend the auto-discovered script to the commands
					newCmd := HookCmd{Run: runStr}
					hookCfg.Commands = append([]HookCmd{newCmd}, hookCfg.Commands...)
					cfg.Hooks[hookName] = hookCfg
				}
			}
		}
	}

	return cfg, nil
}

// CreateDefault creates a starter configuration directory and files based on language choices
func CreateDefault(languages []string) error {
	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return err
	}

	// Always create a pre-commit.sh
	preCommitScript := filepath.Join(ConfigDir, "pre-commit.sh")
	os.WriteFile(preCommitScript, []byte("#!/bin/sh\n# echo \"Running custom pre-commit script\"\n"), 0755)

	cfg := Config{
		Hooks: map[string]HookConfig{
			"pre-commit": {
				Commands: []HookCmd{
					{Run: "./.ultrahooks/pre-commit.sh"},
				},
			},
		},
	}

	preCommitCmds := cfg.Hooks["pre-commit"]

	for _, lang := range languages {
		switch lang {
		case "Go":
			preCommitCmds.Commands = append(preCommitCmds.Commands, 
				HookCmd{Run: "go fmt ./..."},
				HookCmd{Run: "go test ./..."},
			)
		case "Node.js":
			preCommitCmds.Commands = append(preCommitCmds.Commands, 
				HookCmd{Run: "npm run lint"},
				HookCmd{Run: "npm test"},
			)
		case "Python":
			preCommitCmds.Commands = append(preCommitCmds.Commands, 
				HookCmd{Run: "flake8 ."},
				HookCmd{Run: "pytest"},
			)
		case "Rust":
			preCommitCmds.Commands = append(preCommitCmds.Commands, 
				HookCmd{Run: "cargo fmt -- --check"},
				HookCmd{Run: "cargo test"},
			)
		}
	}

	// Update the map reference
	cfg.Hooks["pre-commit"] = preCommitCmds

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(GetConfigPath(), data, 0644)
}
