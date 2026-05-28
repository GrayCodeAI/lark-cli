#!/bin/bash
set -e

# Lark — Multi-agent workspace installer
# Usage: curl -fsSL https://raw.githubusercontent.com/GrayCodeAI/lark-cli/main/install.sh | bash

REPO="GrayCodeAI/lark-cli"
BINARY="lark"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

echo "🐦 Installing Lark..."

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$OS" in
  linux)  PLATFORM="linux" ;;
  darwin) PLATFORM="darwin" ;;
  *)      echo "❌ Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)             echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Get latest release tag
TAG=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" 2>/dev/null | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/' || echo "")

# Fallback to known tag if API fails (private repos)
if [ -z "$TAG" ]; then
  TAG="v0.1.0"
fi

# Try downloading prebuilt binary
URL="https://github.com/$REPO/releases/download/$TAG/${BINARY}-${PLATFORM}-${ARCH}"
echo "📥 Downloading $TAG..."
if curl -fsSL "$URL" -o "$INSTALL_DIR/$BINARY" 2>/dev/null; then
  chmod +x "$INSTALL_DIR/$BINARY"
else
  echo "⚠️  Download failed, building from source..."
  
  # Check prerequisites
  if ! command -v go &>/dev/null; then
    # Try common Go locations
    if [ -x "$HOME/local/go/bin/go" ]; then
      export PATH="$HOME/local/go/bin:$PATH"
    elif [ -x "/usr/local/go/bin/go" ]; then
      export PATH="/usr/local/go/bin:$PATH"
    else
      echo "❌ Go not found. Install from https://go.dev/dl/"
      exit 1
    fi
  fi

  TMPDIR=$(mktemp -d)
  echo "📦 Cloning lark-cli..."
  git clone --depth 1 https://github.com/GrayCodeAI/lark-cli.git "$TMPDIR/lark-cli"
  cd "$TMPDIR/lark-cli"
  
  echo "🔨 Building..."
  go build -o "$BINARY" ./cmd/lark-cli/
  
  echo "📥 Installing to $INSTALL_DIR..."
  mkdir -p "$INSTALL_DIR"
  mv "$BINARY" "$INSTALL_DIR/"
  chmod +x "$INSTALL_DIR/$BINARY"
  
  cd /
  rm -rf "$TMPDIR"
fi

# Verify
if command -v "$BINARY" &>/dev/null; then
  echo ""
  echo "✅ Lark installed successfully!"
  echo ""
  echo "Quick start:"
  echo "  lark init          # Initialize a project"
  echo "  lark connect claude # Connect an agent"
  echo "  lark send general 'Hello from Lark!'"
  echo "  lark status        # Check connection"
  echo ""
  echo "Docs: https://github.com/GrayCodeAI/lark-cli"
else
  echo ""
  echo "⚠️  Installed to $INSTALL_DIR/$BINARY"
  echo "   Add to PATH: export PATH=\"$INSTALL_DIR:\$PATH\""
fi
