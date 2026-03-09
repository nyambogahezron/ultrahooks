package utils

import (
	"os"
	"os/exec"
)

func ExecuteShell(script string, env map[string]string) error {
	cmd := exec.Command("sh", "-c", script)

	// Inject Environment Variables
	cmd.Env = os.Environ() // inherit current env
	for k, v := range env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// ExecuteShellCaptured runs a shell script and captures its output. Useful for parallel execution.
func ExecuteShellCaptured(script string, env map[string]string) (string, error) {
	cmd := exec.Command("sh", "-c", script)

	cmd.Env = os.Environ()
	for k, v := range env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	out, err := cmd.CombinedOutput()
	return string(out), err
}
