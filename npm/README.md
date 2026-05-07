# UltraHooks NPM Wrapper

This is an npm wrapper for the [UltraHooks](https://github.com/nyambogahezron/ultrahooks) CLI tool, which allows users to install and run the pre-compiled Go binary through Node.js/NPM.

## Installation

You can install this globally using `npm`:

```bash
npm install -g ultrahooks
```

Or you can use `npx` to execute it without installing:

```bash
npx ultrahooks --help
```

## How it works

When you install this package, the `postinstall` script runs and automatically downloads the correct pre-built binary for your operating system and architecture (Windows, macOS, or Linux) directly from the GitHub releases page.

## Publishing Updates

When a new version of the `ultrahooks` Go binary is released (e.g., `v1.0.2`):

1. Update the `"version"` field in `npm/package.json` to match the new release (e.g., `"1.0.2"`).
2. Inside the `npm/` directory, run:
   ```bash
   npm publish
   ```

*(Note: Ensure you are logged into npm via `npm login` and have publishing rights for this package).*
