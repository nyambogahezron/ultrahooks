package config

// HookCmd represents a single command to execute within a hook
type HookCmd struct {
	Name string `yaml:"name"`
	Run  string `yaml:"run"`
}

// Config represents the .ultrahooks.yml structure
type Config struct {
	Hooks map[string][]HookCmd `yaml:"hooks"`
}
