#!/usr/bin/env bash
set -e

REPO="nyambogahezron/ultrahooks"
GITHUB_URL="https://github.com/$REPO/releases/latest/download"

echo "=========================================="
echo "      Installing UltraHooks (Remote)      "
echo "=========================================="

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  linux*)               OS="linux" ;;
  darwin*)              OS="darwin" ;;
  mingw*|msys*|cygwin*) OS="windows" ;;
  *)                    echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Detect Architecture
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)            echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

BINARY_NAME="ultrahooks_${OS}_${ARCH}"
if [ "$OS" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

DOWNLOAD_URL="${GITHUB_URL}/${BINARY_NAME}"

# Install Directory
if [ "$OS" = "windows" ]; then
    INSTALL_DIR="$HOME/bin"
    SUDO=""
else
    INSTALL_DIR="/usr/local/bin"
    if [ ! -w "$INSTALL_DIR" ]; then
        SUDO="sudo "
        echo "Requires sudo privileges to install to $INSTALL_DIR"
    else
        SUDO=""
    fi
fi

mkdir -p "$INSTALL_DIR"

echo "Downloading latest release for $OS ($ARCH)..."
echo "Fetching from: $DOWNLOAD_URL"

TMP_DIR=$(mktemp -d)
if [ "$OS" = "windows" ]; then
    DEST_PATH="$TMP_DIR/ultrahooks.exe"
else
    DEST_PATH="$TMP_DIR/ultrahooks"
fi

curl -sL --fail -o "$DEST_PATH" "$DOWNLOAD_URL" || {
    echo "Error: Failed to download the binary. Please check if the release exists on GitHub."
    exit 1
}

chmod +x "$DEST_PATH"

echo "Installing to $INSTALL_DIR..."
if [ "$OS" = "windows" ]; then
    mv "$DEST_PATH" "$INSTALL_DIR/ultrahooks.exe"
else
    $SUDO mv "$DEST_PATH" "$INSTALL_DIR/ultrahooks"
fi
rm -rf "$TMP_DIR"

echo "Setting up shell auto-completion..."
SHELL_NAME=$(basename "$SHELL")
if [ "$SHELL_NAME" = "zsh" ]; then
    echo "Adding zsh completion..."
    ultrahooks completion zsh > ~/.ultrahooks_completion.zsh
    if ! grep -q ".ultrahooks_completion.zsh" ~/.zshrc; then
        echo "source ~/.ultrahooks_completion.zsh" >> ~/.zshrc
    fi
elif [ "$SHELL_NAME" = "bash" ]; then
    echo "Adding bash completion..."
    if [ -d "/usr/share/bash-completion/completions" ] && [ "$OS" != "windows" ]; then
        $SUDO sh -c 'ultrahooks completion bash > /usr/share/bash-completion/completions/ultrahooks'
    elif [ -d "/etc/bash_completion.d" ] && [ "$OS" != "windows" ]; then
        $SUDO sh -c 'ultrahooks completion bash > /etc/bash_completion.d/ultrahooks'
    else
        ultrahooks completion bash > ~/.ultrahooks_completion.bash
        if ! grep -q ".ultrahooks_completion.bash" ~/.bashrc; then
            echo "source ~/.ultrahooks_completion.bash" >> ~/.bashrc
        fi
    fi
fi

echo "=========================================="
echo " UltraHooks Installed Successfully! 🎉    "
echo " Make sure $INSTALL_DIR is in your PATH.  "
echo " Run 'ultrahooks --help' to get started.  "
echo "=========================================="
