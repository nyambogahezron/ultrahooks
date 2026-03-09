# CLI Commands Reference

UltraHooks utilizes a simple, straightforward CLI mapped intuitively to your workflow.

---

### `ultrahooks init`
**Description:** Initializes the UltraHooks context in your current Git repository.
**Behavior:**
1. Triggers an interactive CLI prompt using `survey`.
2. Asks: `"What languages does your project use? [Go, Node.js, Python, Rust, Other]"`
3. Scaffolds the `.ultrahooks/` directory.
4. Drops a standard `.ultrahooks/pre-commit.sh` shell script template.
5. Dynamically populates `.ultrahooks/config.yaml` with commands matching your selected languages (e.g. `npm run lint`, `go fmt`).

---

### `ultrahooks add <hookName...>`
**Description:** Easily scaffolds new template files for specific Git hooks and simultaneously wires them up.
**Behavior:**
- Generates `.ultrahooks/<hookName>.sh` and applies `chmod +x`.
- Immediately installs the Git proxy script into `.git/hooks/<hookName>`.
- Multi-argument friendly.
- Integrates shell auto-completion for Git standard hook names to prevent typos!

**Example:**
```bash
ultrahooks add pre-commit pre-push
```

---

### `ultrahooks install`
**Description:** Synchronizes your `.git/hooks/` folder to match your active UltraHooks configuration.
**Behavior:**
- Reads your `.ultrahooks/config.yaml` and scans the `.ultrahooks/` directory.
- For every active hook discovered, it writes a proxy execution script into `.git/hooks/`.
- **Note:** It does *not* blindly install all 28 Git hooks. It is precise and surgical, only installing exactly what you are using.

---

### `ultrahooks run <hookName>`
**Description:** Executes the logic for a specific hook. (This is primarily what the Git proxy scripts call).
**Behavior:**
- Loads `config.yaml` map for `<hookName>`.
- Auto-discovers any shell scripts (`.sh`, `.bat`, `.ps1`) in the `.ultrahooks/` directory matching the name and prepends them.
- Executes them sequentially (or concurrently if `parallel: true` is set).

---

### `ultrahooks remove [hookName...] [--all]`
**Description:** Cleanly unwires your Git hooks and deletes scaffolding.
**Behavior:**
- **Specific Removal**: `ultrahooks remove pre-commit` will unwire `.git/hooks/pre-commit` and delete `.ultrahooks/pre-commit.sh`, leaving `pre-push` perfectly intact.
- **Global Wipe**: `ultrahooks remove --all` scans `.git/hooks/` and finds *any* proxy script managed by UltraHooks, deleting them. It then completely wipes all recognized shell scripts in your `.ultrahooks/` directory.

---

### `ultrahooks doctor`
**Description:** A health-check diagnostic scanner.
**Behavior:**
- Checks if you are in a Git repository.
- Validates the syntactical health of `.ultrahooks/config.yaml`.
- Analyzes custom script permissions (`chmod +x` executable bit).
- Verifies that proxy links inside `.git/hooks/` actually exist for the hooks you have defined.
