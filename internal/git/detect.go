package git

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// GetStagedFiles returns a list of files currently staged in Git
func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only", "--diff-filter=ACM")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rawFiles := strings.Split(string(out), "\n")
	var files []string
	for _, f := range rawFiles {
		f = strings.TrimSpace(f)
		if f != "" {
			files = append(files, f)
		}
	}
	return files, nil
}
