# Contributing to UltraHooks

First off, thank you for considering contributing to UltraHooks! It’s people like you who make the open-source community such an incredible place to learn, inspire, and create.

This document serves as a guideline for contributing efficiently. 

---

## 1. Setup Your Development Environment

UltraHooks is written purely in Go.

### Prerequisites:
- **Go 1.21+** installed.
- Git installed.

### Steps to Compile Locally:
1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/ultrahooks.git
   cd ultrahooks
   ```
3. **Install Dependencies** (we use very few external packages):
   ```bash
   go mod download
   ```
4. **Compile the binary**:
   ```bash
   go build -o ultrahooks
   ```

You can now use `./ultrahooks` locally to test your changes.

---

## 2. Making Changes

- Before writing massive new features, please open an **Issue** on GitHub to discuss the architectural approach. UltraHooks aims to be extremely lightweight, and we try to avoid adding bloated features unless strongly justified.
- Create a new branch for your feature/bugfix:
  ```bash
  git checkout -b feat/my-new-feature
  ```
- Write clear, concise Go code following standard `gofmt` paradigms.

---

## 3. Submitting a Pull Request

When submitting a PR, ensure that:
1. You have run `go fmt ./...` on your codebase to format correctly.
2. You have executed `go test ./...` and no regressions exist.
3. Your commit messages clearly explain the "Why" and "What" of your changes. Include issue numbers (e.g., `fixes #12`).

---

## 4. Code of Conduct

All contributors are expected to abide by our Code of Conduct:
- Be respectful and constructive during PR reviews.
- Focus on the technical merit of the code, not the person authoring it.
- Harassment of any kind is strictly prohibited.
