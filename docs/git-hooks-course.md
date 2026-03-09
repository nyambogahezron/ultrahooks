# Git Hooks Masterclass

Git hooks are custom scripts that Git executes before or after events such as: commit, push, merge, and receive. 

UltraHooks allows you to harness these events effortlessly. This guide covers the most important Git hooks, when they fire, and why you should use them.

---

## 1. `pre-commit` (The Most Important Hook)
**When it fires:** Immediately after you type `git commit` but *before* Git even prompts you for a commit message or creates the commit object.
**Why use it:** To catch code issues early. If this hook exits with a failure (non-zero status), the commit is aborted. This prevents broken styling or failing tests from ever entering the Git timeline.

**Common Use Cases:**
- Code Formatters (`go fmt`, `prettier`, `black`)
- Linters (`eslint`, `golangci-lint`, `flake8`)
- Basic, rapid unit tests
- Type checkers (`tsc --noEmit`)

**UltraHooks Example (`.ultrahooks/config.yaml`):**
```yaml
hooks:
  pre-commit:
    commands:
      - name: Lint JSON and JS
        run: npm run lint
```

---

## 2. `prepare-commit-msg`
**When it fires:** Before the commit message editor is fired up, but after the default message is created. It takes 1 to 3 parameters (the name of the file that holds the commit message, the source of the message, and the commit SHA-1).
**Why use it:** To automatically inject data into your commit messages.

**Common Use Cases:**
- Injecting Jira/Linear issue tracking IDs dynamically into the commit subject based on your branch name (e.g. `feat/JIRA-123` -> `[JIRA-123] Added new button`).
- Enforcing Conventional Commits format templates.

---

## 3. `commit-msg`
**When it fires:** After you save your commit message and exit the editor, but before the commit is permanently finalized. It takes one parameter: the path to a temporary file containing the commit message you just wrote.
**Why use it:** To validate the quality or enforce the rules of your commit message. If it exits non-zero, the commit is aborted.

**Common Use Cases:**
- Validating that the message follows a strict pattern using tools like `commitlint`.
- Rejecting empty strings or messages under 10 characters.
- Ensuring a ticket ID is present.

---

## 4. `post-commit`
**When it fires:** Immediately after the entire commit process is completed successfully.
**Why use it:** Primarily for notification or logging purposes. It cannot affect the outcome of the `git commit` command.

**Common Use Cases:**
- Notifying a Slack or Discord channel.
- Triggering a quick local desktop notification alerting you that a big refactor commit finally compiled and saved.

---

## 5. `pre-push`
**When it fires:** After you type `git push` but *before* any objects have been transferred to the remote (origin).
**Why use it:** To run the heavy, long-running validations that are too slow to run on every single commit, saving you from pushing broken code that would instantly fail the CI/CD pipeline on GitHub/GitLab.

**Common Use Cases:**
- Running the full, exhaustive E2E (End-to-End) Cypress or Playwright test suite.
- Checking integration tests against a test database.
- Checking code coverage percentages.

**UltraHooks Example:**
```yaml
hooks:
  pre-push:
    commands:
      - name: E2E Tests
        run: npm run test:e2e
```

---

## 6. `post-checkout`
**When it fires:** After you run a successful `git checkout` (e.g. switching branches).
**Why use it:** To ensure the state of your development environment matches the branch you just landed on.

**Common Use Cases:**
- Automatically running `npm ci` or `go mod download` if the `package.json` or `go.mod` changed between branches, so you aren't fighting missing dependencies.
- Resetting a local development database schema to match the branch's migration state.

---

## 7. `pre-rebase`
**When it fires:** Before you rebase a branch.
**Why use it:** Primarily utilized to prevent developers from accidentally rebasing commits that have already been pushed to a shared remote repository, which rewrites history and causes chaos for teammates.

---

## Additional Resources
- [Official Git Hooks Documentation](https://git-scm.com/docs/githooks)
- [Atlassian Git Hooks Tutorial](https://www.atlassian.com/git/tutorials/macgit-hooks)
- [Conventional Commits Specification](https://www.conventionalcommits.org/en/v1.0.0/)
