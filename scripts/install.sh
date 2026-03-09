# local installer

#!/usr/bin/env bash
set -e

echo "Installing UltraHooks (Local)..."

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  linux*)               OS="linux" ;;
  darwin*)              OS="darwin" ;;
  mingw*|msys*|cygwin*) OS="windows" ;;
  *)                    echo "Unsupported OS: $OS"; exit 1 ;;
esac

echo "Building UltraHooks for $OS..."

# Install Directory
if [ "$OS" = "windows" ]; then
    INSTALL_DIR="$HOME/bin"
    go build -o ultrahooks.exe main.go
    DEST="ultrahooks.exe"
else
    INSTALL_DIR="$HOME/.local/bin"
    go build -o ultrahooks main.go
    DEST="ultrahooks"
fi

echo "Installing to $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
mv "$DEST" "$INSTALL_DIR/"

echo "UltraHooks Built & Installed! 🎉"
echo "Make sure $INSTALL_DIR is in your PATH."
echo "Run 'ultrahooks --help' to get started."
