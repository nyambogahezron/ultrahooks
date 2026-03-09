package config

// HookCmd represents a single command to execute within a hook
type HookCmd struct {
	Name     string            `yaml:"name"`
	Run      string            `yaml:"run"`
	Parallel bool              `yaml:"parallel,omitempty"`
	Env      map[string]string `yaml:"env,omitempty"`
}

// HookConfig allows configuring the hook execution overall OR just a list of commands
type HookConfig struct {
	Parallel bool      `yaml:"parallel,omitempty"`
	Commands []HookCmd `yaml:"commands,omitempty"`
}

// Config represents the .ultrahooks.yml structure
type Config struct {
	Hooks map[string]HookConfig `yaml:"hooks"`
}

// UnmarshalYAML implements custom unmarshaling to support both:
// hooks:
//
//	pre-commit:
//	  - name: Foo
//	    run: bar
//
// AND
// hooks:
//
//	pre-commit:
//	  parallel: true
//	  commands:
//	    - name: Foo
//	      run: bar
func (hc *HookConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try to parse it as just a list of commands (the V1/V2 backward compatible way)
	var commands []HookCmd
	if err := unmarshal(&commands); err == nil {
		hc.Commands = commands
		return nil
	}

	// If that fails, parse it as the extended HookConfig object
	type extendedHookConfig HookConfig
	var ehc extendedHookConfig
	if err := unmarshal(&ehc); err != nil {
		return err
	}

	*hc = HookConfig(ehc)
	return nil
}
