package config

import "gopkg.in/yaml.v3"

type HookCmd struct {
	Name     string            `yaml:"name"`
	Run      string            `yaml:"run"`
	Parallel bool              `yaml:"parallel,omitempty"`
	Env      map[string]string `yaml:"env,omitempty"`
}

type HookConfig struct {
	Parallel bool      `yaml:"parallel,omitempty"`
	Commands []HookCmd `yaml:"commands,omitempty"`
}

type Config struct {
	Hooks map[string]HookConfig `yaml:"hooks"`
}

func (hc *HookConfig) UnmarshalYAML(value *yaml.Node) error {
	// First try to parse it as just a list of commands (the V1/V2 backward compatible way)
	if value.Kind == yaml.SequenceNode {
		var commands []HookCmd
		if err := value.Decode(&commands); err != nil {
			return err
		}
		hc.Commands = commands
		return nil
	}

	// If that fails, parse it as the extended HookConfig object
	type extendedHookConfig HookConfig
	var ehc extendedHookConfig
	if err := value.Decode(&ehc); err != nil {
		return err
	}

	*hc = HookConfig(ehc)
	return nil
}
