# UltraHooks

UltraHooks is a fast, lightweight, universal Git hooks manager written in Go.
It is similar to Husky and Lefthook but simpler, faster, and language-agnostic.

UltraHooks works with ANY programming language (Go, Node.js, Python, Rust, Java, PHP, C++, etc) and does NOT depend on Node.js or Python.
It is distributed as a single Go binary.

## Features

- **Language Agnostic**: Works seamlessly with any language.
- **Fast**: Written in Go, minimal overhead.
- **Simple Configuration**: Easy to understand YAML config `.ultrahooks/config.yaml`.
- **Custom Scripts**: Automatically runs `<hook>.sh` shell scripts from `.ultrahooks/`.
- **Zero Dependencies**: You only need the UltraHooks binary.
- **Live Output Streaming**: See the execution of your hooks live.
- **Fail-Fast**: Stops on the first error to save your time.

## Installation

### Method 1: Go Install
```bash
go install github.com/nyambogahezron/ultrahooks@latest
```

### Method 2: Install Script (Local/Cloned Repo)
If you've cloned the repository and want to build and install it locally to your machine (`~/.local/bin`):
```bash
git clone https://github.com/nyambogahezron/ultrahooks.git
cd ultrahooks
./install.sh
```

## Usage

1. **Initialize UltraHooks in your repository:**
   ```bash
   ultrahooks init
   ```
   This creates a default `.ultrahooks` configuration directory.

2. **Add a new hook:**
   ```bash
   ultrahooks add pre-commit
   ```
   This scaffolds a new `.ultrahooks/pre-commit.sh` template and automatically installs the Git proxy script for it!

3. **Configure your hooks:**
   Edit the `.ultrahooks/config.yaml` file to run the commands you want, or edit the shell scripts like `.ultrahooks/pre-commit.sh` directly.

## Example Config

`.ultrahooks/config.yaml`:
```yaml
hooks:
  pre-commit:
    - run: ./.ultrahooks/pre-commit.sh
    - run: npm run lint
    - run: go test ./...

  pre-push:
    - run: ./.ultrahooks/pre-push.sh
    - run: go build ./...
```

## Commands

- `ultrahooks init` - Initialize hooks system
- `ultrahooks add <hook>` - Scaffold a new `.ultrahooks/<hook>.sh` template and instantly wire it to Git.
- `ultrahooks install` - Re-install proxy scripts into git for the active hooks.
- `ultrahooks run <hook>` - Run hooks for a specific event
- `ultrahooks remove [hooks...]` - Unwire specific Git hooks and delete their custom scripts.
- `ultrahooks remove --all` - Fully unwire all Git hooks managed by UltraHooks and delete all custom scripts.
