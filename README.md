<p align="center">
  <img src="docs/public/logo.png" alt="UltraHooks Logo" width="200"/>
</p>

<p align="center">
 <h1>UltraHooks - The Universal Git Hooks Manager</h1>
</p>

<p align="center">

[![License-MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go-Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/dl/)
[![Cross-Platform](https://img.shields.io/badge/Platform-Windows%20|%20macOS%20|%20Linux-lightgrey.svg)]()

</p>

**UltraHooks** is a blazingly fast, lightweight, and language-agnostic Git hooks manager written in Go. It is designed to be a simpler, faster alternative to tools like Husky and Lefthook, without requiring Node.js, Python, or Ruby dependencies. Built with performance as a top priority, UltraHooks provides parallel execution, file change detection, and cross-platform native execution.

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Documentation](#documentation)
- [Contributing](#contributing)

# Features

- **Blazingly Fast**: Written in Go with minimal overhead. Supports concurrent parallel hook execution via goroutines (`parallel: true`).
- **Language Agnostic**: Works perfectly out of the box for Go, Node.js, Python, Rust, PHP, C++, and more. No `node_modules`.
- **File Change Detection**: Native support for the `{staged_files}` keyword, automatically skipping execution if files are untouched.
- **Cross-Platform Native**: Executes flawlessly on macOS, Linux, and Windows (natively discovering and picking up `.sh`, `.bat`, or `.ps1` files).
- **Extremely Clean**: Surgical precision installation directly wires hooks from your shared repository configurations without polluting your root `.git/hooks` directory.

# Installation

```sh
# Clone the repository
git clone https://github.com/nyambogahezron/ultrahooks.git
cd ultrahooks

# Build and install the application locally
./local_install.sh
```

## Global Installation (Recommended)

To install the latest compiled binary of UltraHooks straight to your path and automatically setup shell-completions (macOS, Linux, and Windows via Git Bash):

```bash
curl -sSL https://raw.githubusercontent.com/nyambogahezron/ultrahooks/main/install.sh | bash
```

# Usage

```sh
ultrahooks COMMAND
# runs the command
ultrahooks help COMMAND
# outputs help for specific command
```

# Commands

<!-- commands -->
* [`ultrahooks init`](#ultrahooks-init)
* [`ultrahooks add`](#ultrahooks-add)
* [`ultrahooks install`](#ultrahooks-install)
* [`ultrahooks run`](#ultrahooks-run)
* [`ultrahooks remove`](#ultrahooks-remove)
* [`ultrahooks doctor`](#ultrahooks-doctor)

## `ultrahooks init`

Initializes the UltraHooks context in your current Git repository.

```
USAGE
  $ ultrahooks init

DESCRIPTION
  Triggers an interactive CLI prompt to scaffold the .ultrahooks/ directory and dynamically generate a tailored config.yaml.
```

## `ultrahooks add`

Easily scaffolds new template files for specific Git hooks and simultaneously wires them up.

```
USAGE
  $ ultrahooks add <hookName...>

DESCRIPTION
  Generates .ultrahooks/<hook>.sh and immediately installs the Git proxy script into .git/hooks/<hook>. Multi-argument friendly.
```

## `ultrahooks install`

Synchronizes your `.git/hooks/` folder to match your active UltraHooks configuration.

```
USAGE
  $ ultrahooks install

DESCRIPTION
  Reads your .ultrahooks/config.yaml and writes safe proxy execution scripts into .git/hooks/ only for the hooks you explicitly use.
```

## `ultrahooks run`

Executes the logic for a specific hook. (This is primarily what the Git proxy scripts call).

```
USAGE
  $ ultrahooks run <hookName>

DESCRIPTION
  Reads config.yaml and auto-discovers shell scripts, executing them sequentially or concurrently.
```

## `ultrahooks remove`

Cleanly unwires your Git hooks and deletes scaffolding.

```
USAGE
  $ ultrahooks remove [hookName...] [--all]

DESCRIPTION
  Unwires specific hooks or use --all to completely wipe any trace of UltraHooks proxy scripts from the repository.
```

## `ultrahooks doctor`

A health-check diagnostic scanner.

```
USAGE
  $ ultrahooks doctor

DESCRIPTION
  Scans your repository for missing directories, invalid configurations, non-executable scripts, and missing Git proxy links.
```

# Documentation

- **Documentation**: Everything is inside the `/docs` directory. (VitePress site available)
- **Features Deep Dive**: [features.md](docs/features.md)
- **Architecture**: [architecture.md](docs/architecture.md)
- **Git Hooks Course**: [git-hooks-course.md](docs/git-hooks-course.md)

# Contributing

We welcome contributions from the community! 

- Read our [Contributing Guidelines](CONTRIBUTING.md) to get started
- Check our [Code of Conduct](CODE_OF_CONDUCT.md)
- Browse [open issues](https://github.com/nyambogahezron/ultrahooks/issues) or create a new one

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Security

Security is paramount for UltraHooks as it handles local shell execution. If you discover a security vulnerability, please follow our [Security Policy](SECURITY.md) for responsible disclosure. **Do not report security vulnerabilities through public GitHub issues.**
