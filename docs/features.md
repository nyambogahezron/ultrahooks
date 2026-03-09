# Features & Capabilities

UltraHooks isn't just a shell script runner. It contains powerful internal scheduling and execution mechanics to make your CI/CD pipelines incredibly fast directly on your local machine.

---

## 1. Parallel Execution (`parallel: true`)

By default, commands execute sequentially. If you have a linter, a formatter, and a unit test suite, running them sequentially during a `pre-commit` might take 10-15 seconds.

With UltraHooks, you can run them **concurrently** using Go routines by modifying your `.ultrahooks/config.yaml`:

```yaml
hooks:
  pre-commit:
    parallel: true
    commands:
      - name: Linter
        run: npm run lint
      - name: Formatter
        run: go fmt ./...
      - name: Unit Tests
        run: go test ./...
```

**How Output is Handled:**
UltraHooks safely captures the `stdout` and `stderr` buffers for each routine. It waits until all routines finish, and then prints the outputs sequentially. This guarantees that your terminal output is never garbled with mixed lines from NodeJS and Go!

---

## 2. File Change Detection (`{staged_files}`)

It is inefficient to run `eslint .` or `go test ./...` on your *entire* codebase if you only modified a single Markdown file.

UltraHooks features an internal Git traversal engine that can resolve the exact files you've staged for commit using `git diff --cached`. To use it, simply inject the `{staged_files}` token into your `run` command string:

```yaml
hooks:
  pre-commit:
    commands:
      - name: Lint Modified Files
        run: eslint {staged_files}
```

**Smart Skipping:** 
If a command explicitly requires `{staged_files}`, but you haven't staged any files that match your filtering, UltraHooks will **instantly skip the command entirely** to save you time.

---

## 3. Environment Variable Injection (`env`)

Sometimes you need to trick a command into running in a specific mode (e.g., `NODE_ENV=test` or `CI=true`) but you don't want to export that globally in your shell or pollute your generic `.sh` files.

You can inject environment maps on a per-command basis:

```yaml
hooks:
  pre-commit:
    commands:
      - name: Build
        run: npm run build
        env:
          NODE_ENV: "production"
          DISABLE_WARNINGS: "1"
```

---

## 4. Cross-Platform Native Discovery

True language-agnostic tools must be Operating System agnostic. 

While `.sh` scripts work natively on Linux and macOS (and often Git Bash on Windows), UltraHooks natively scans and discovers `.bat` and `.ps1` (PowerShell) scripts inside the `.ultrahooks/` directory alongside `.sh` scripts. 

Simply name your file `pre-commit.bat` or `pre-commit.ps1`, and UltraHooks will execute it seamlessly on Windows without requiring special Docker or Bash emulation layers.

---

## 5. Doctor Health Scanner

If something goes wrong (e.g., a hook isn't firing, or a script is failing to execute), simply run:

```bash
ultrahooks doctor
```

This acts as a diagnostic AI that scans:
- If you are actually inside a `.git` repository folder tree.
- If your `config.yaml` has valid YAML syntax.
- If your custom `.sh` scripts have the correct execution bits (`chmod +x`).
- If your Git hooks in `.git/hooks/` are correctly wired to the UltraHooks binary.
