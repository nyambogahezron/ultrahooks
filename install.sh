#!/usr/bin/env bash

set -e

echo "Building UltraHooks..."
go build -o ultrahooks main.go

INSTALL_DIR="$HOME/.local/bin"

echo "Installing to $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
mv ultrahooks "$INSTALL_DIR/"

echo "✔ UltraHooks installed successfully to $INSTALL_DIR/ultrahooks"
echo ""
echo "Note: Make sure $INSTALL_DIR is in your system's PATH."
