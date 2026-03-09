# Installation Guidelines

You can install UltraHooks globally on your system, or compile it from source for your specific repository.

---

## Option 1: Global Installation via script (Recommended)

To install the latest compiled binary of UltraHooks straight to your path, run our `install.sh` script.

This script will automatically copy the `ultrahooks` binary into `/usr/local/bin` (or your local user bin path) so it can be accessed globally from any directory.

```bash
curl -sSL https://raw.githubusercontent.com/nyambogahezron/ultrahooks/main/install.sh | bash
```

**Verify Installation:**
```bash
ultrahooks --version
```
*(If the command is not found, ensure `/usr/local/bin` is in your `$PATH`.)*

---

## Option 2: Compile from Source (Go Developers)

If you have Go installed on your machine (`1.21+`), you can globally install the binary directly from GitHub using `go install`:

```bash
go install github.com/nyambogahezron/ultrahooks@latest
```

Ensure your Go `bin/` directory (usually `~/go/bin`) is included in your `$PATH`.

---

## Option 3: Local Repository Building

If you have cloned this repository and wish to build it manually for local testing or contribution:

1. Clone the repository:
   ```bash
   git clone https://github.com/nyambogahezron/ultrahooks
   cd ultrahooks
   ```
2. Build the binary outputs:
   ```bash
   go build -o ultrahooks
   ```
3. Move it locally:
   ```bash
   sudo cp ./ultrahooks /usr/local/bin/ultrahooks
   ```
