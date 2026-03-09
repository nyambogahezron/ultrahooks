# UltraHooks Architecture

UltraHooks is designed to be as minimally invasive as possible. It avoids locking you into proprietary systems or placing global dependencies on your developers' machines. 

Here is how the architecture is wired together:

## 1. The Core Repository Separation

UltraHooks separates its logic into two distinct areas:
1. **`.git/hooks/`**: The standard Git directory.
2. **`.ultrahooks/`**: The custom project directory that developers actually check into source control.

---

## 2. The Git Proxy Scripts (`.git/hooks/`)

By standard Git design, when an event like `git commit` fires, Git automatically looks for an executable file named `.git/hooks/pre-commit`. 

When you run `ultrahooks install` or `ultrahooks add pre-commit`, UltraHooks generates a **Proxy Script** inside that directory. The proxy script looks like this:

```sh
#!/bin/sh
# ultrahooks run
ultrahooks run pre-commit "$@"
```

**Why a Proxy?**
We do not place your actual custom bash logic (like `npm run lint`) inside `.git/hooks/`. Why? Because the `.git` folder is not tracked by Git! If you put your logic there, your team members wouldn't get the updates when they pull from the repository.

Instead, the proxy script simply intercepts the Git event and hands it off to the UltraHooks Go binary.

---

## 3. The Custom Logic Directory (`.ultrahooks/`)

When the Proxy Script executes `ultrahooks run pre-commit`, the UltraHooks binary wakes up and looks inside your project's tracked `.ultrahooks/` directory.

It does two things in order:
1. **Script Auto-Discovery**: It looks for any script named `pre-commit.sh`, `pre-commit.bat`, or `pre-commit.ps1`. If it finds one, it executes it natively.
2. **YAML Pipeline Execution**: It parses `.ultrahooks/config.yaml`. If it finds a `pre-commit` map, it executes all the YAML-defined commands sequentially (or in parallel if `parallel: true` is set).

---

## 4. Why is this Architecture Powerful?

- **Shareable**: The `.ultrahooks/` directory is committed to source control. When a new developer clones the repository, they get all the hooks, scripts, and configurations instantly. All they have to do is run `ultrahooks install` to wire their local `.git` folder up.
- **Language Agnostic**: Because the binary is compiled Go, it doesn't care if the underlying command is `go test`, `npm run build`, `cargo fmt`, or `docker-compose up`. It just passes the command to the OS shell.
- **Zero Global Garbage**: UltraHooks cleans up after itself perfectly using `ultrahooks remove --all`, which scans for the `ultrahooks run` signature in proxy scripts and deletes them cleanly without touching other Git hooks you might have installed manually.
