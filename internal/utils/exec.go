package utils

import (
	"os"
	"os/exec"
)

func ExecuteShell(script string) error {
	cmd := exec.Command("sh", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
