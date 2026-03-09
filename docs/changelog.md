# Changelog

All notable changes to the UltraHooks project will be documented in this file.

## [v1.0.0] - 2026-03-09

**V5 Release: Advanced Features**
- **Parallel Execution:** Concurrent execution of hook commands via `goroutines`, natively buffered to prevent garbled output (`parallel: true`).
- **File Change Detection:** Native support for the `{staged_files}` keyword utilizing Git diff to dynamically skip execution if files are untouched.
- **Environment Injections:** Pass arbitrary `env` maps directly to commands in the `config.yaml` without global pollution.
- **Windows Compatability:** Full support for discovering native `.bat` and `.ps1` (PowerShell) custom templates natively alongside `.sh` scripts.
- **Interactive Initializer:** Re-engineered `ultrahooks init` with Survey library integration to generate perfectly tailored YAML configurations depending on the user's stack (Go, Node.js, Python, Rust).
- **Diagnostics:** Introduced `ultrahooks doctor` to analyze YAML syntax, file permissions, and Git proxy linkages to rapidly solve configuration issues.

**V4 Release: Remove & Clean-Up**
- Implemented `ultrahooks remove [hook...]` to dynamically unwire active git proxy scripts and delete custom custom `.ultrahooks/` `.sh` scaffolding templates concurrently.
- Implemented `ultrahooks remove --all` for a 1-click total cleanup of all generated proxy configurations and scripts, reverting a repository back to its clean state.

**V3 Release: Precision Core**
- Introduced surgical `ultrahooks install` mechanisms that discover active configurations and *only* wire up `.git/hooks/` files for hooks actively in use, keeping the Git directory pristinely clean.
- Added `ultrahooks add <hook>` to generate templates and instantly install Git proxies simultaneously with full shell-completion support.

**V2 Release: The `.ultrahooks` Architecture**
- Migrated out of the restrictive `.ultrahooks.yml` architecture format and into the rich `.ultrahooks` configuration directory format.
- Added support to auto-discover executable script files natively without requiring explicit definitions inside `config.yaml`.

**V1 Release: The Foundation**
- Foundation established in Go.
- `init`, `install`, `uninstall`, and `run` command structure solidified.
- Custom logging, colorized formatting, and streaming execution logic defined.
