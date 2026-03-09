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
		Hooks: make(map[string][]HookCmd),
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
		cfg.Hooks = make(map[string][]HookCmd)
	}

	// Auto-discover shell scripts in the config directory to make config.yaml optional
	if entries, err := os.ReadDir(ConfigDir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".sh" {
				hookName := entry.Name()[:len(entry.Name())-3] // Remove .sh
				scriptPath := filepath.Join(ConfigDir, entry.Name())
				runStr := "./" + scriptPath

				cmds := cfg.Hooks[hookName]
				found := false
				for _, cmd := range cmds {
					if cmd.Run == runStr {
						found = true
						break
					}
				}

				if !found {
					// Prepend the auto-discovered script to the commands
					newCmd := HookCmd{Run: runStr}
					cfg.Hooks[hookName] = append([]HookCmd{newCmd}, cmds...)
				}
			}
		}
	}

	return cfg, nil
}

// CreateDefault creates a starter configuration directory and files
func CreateDefault() error {
	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return err
	}

	// Create some default executable scripts as examples
	preCommitScript := filepath.Join(ConfigDir, "pre-commit.sh")
	os.WriteFile(preCommitScript, []byte("##!/bin/sh\n# echo \"Running custom pre-commit script\"\n"), 0755)

	prePushScript := filepath.Join(ConfigDir, "pre-push.sh")
	os.WriteFile(prePushScript, []byte("##!/bin/sh\n# echo \"Running custom pre-push script\"\n"), 0755)

	cfg := Config{
		Hooks: map[string][]HookCmd{
			"pre-commit": {
				{Run: "./.ultrahooks/pre-commit.sh"},
				{Run: "go fmt ./..."},
				{Run: "go test ./..."},
			},
			"pre-push": {
				{Run: "./.ultrahooks/pre-push.sh"},
				{Run: "go build ./..."},
			},
		},
	}

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(GetConfigPath(), data, 0644)
}
