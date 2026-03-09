# UltraHooks Documentation

Welcome to the official documentation for **UltraHooks**! 

UltraHooks is a blazingly fast, lightweight, and language-agnostic Git hooks manager written in Go. It is designed to be a simpler, faster alternative to tools like Husky and Lefthook, without requiring Node.js, Python, or Ruby dependencies.

## Table of Contents

1. [Features & Capabilities](./features.md)
   Learn about Parallel Execution, Staged File Detection, and Environment Injection.
2. [Installation Guide](./installation.md)
   How to install UltraHooks locally or via our remote `curl` script.
3. [CLI Commands Reference](./commands.md)
   Detailed usage for `init`, `add`, `install`, `run`, `remove`, and `doctor`.
4. [Architecture & How It Works](./architecture.md)
   A deep dive into how UltraHooks wires itself cleanly into `.git/hooks/` and discovers your `.ultrahooks/` shell scripts.
5. [Git Hooks Masterclass](./git-hooks-course.md)
   A comprehensive guide explaining what every Git Hook does and how you can use them effectively in your daily workflow.
6. [Contributing & Collaboration](./contributing.md)
   How to build from source, run tests, and contribute back to the open-source repository.
7. [Changelog](./changelog.md)
   Historical tracking of new features and major version upgrades.

---

### Why UltraHooks?
- **Zero Dependencies**: It's a single Go binary. No `node_modules`.
- **Incredibly Fast**: Boot time is virtually zero. Supports concurrent goroutine execution of hooks.
- **Language Agnostic**: Works perfectly out of the box for Go, Node.js, Python, Rust, PHP, C++, and more.
- **Cross-Platform**: Executes flawlessly on macOS, Linux, and Windows (natively picking up `.sh`, `.bat`, or `.ps1` files).
