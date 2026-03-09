package git

import (
	"errors"
	"os"
	"path/filepath"
)

func FindGitDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		gitPath := filepath.Join(dir, ".git")
		if info, err := os.Stat(gitPath); err == nil && info.IsDir() {
			return gitPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", errors.New("not a git repository (or any of the parent directories)")
}
