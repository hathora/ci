#!/usr/bin/env bash

set -e

REPO_OWNER="hathora"
REPO_NAME="ci"
BINARY_NAME="hathora-ci"

# Get the latest release tag from GitHub
LATEST_RELEASE=$(curl -s https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest | grep "tag_name" | cut -d '"' -f 4)

# Detect the operating system and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux*)
    case "$ARCH" in
      x86_64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-linux-amd64"
        ;;
      aarch64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-linux-arm64"
        ;;
      *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac
    INSTALL_DIR="/usr/local/bin"
    ;;
  darwin*)
    case "$ARCH" in
      x86_64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-darwin-amd64"
        ;;
      arm64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-darwin-arm64"
        ;;
      *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac
    INSTALL_DIR="/usr/local/bin"
    ;;
  msys*|cygwin*|mingw*)
    case "$ARCH" in
      x86_64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-windows-amd64"
        ;;
      arm64)
        DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_RELEASE}/${BINARY_NAME}-windows-arm64"
        ;;
      *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
    esac
    INSTALL_DIR="${USERPROFILE}\\AppData\\Local\\Microsoft\\WindowsApps"
    ;;
  *)
    echo "Unsupported operating system: $OS"
    exit 1
    ;;
esac

# Download the binary
echo "Downloading ${BINARY_NAME} version ${LATEST_RELEASE}..."
curl -s -L -o "${BINARY_NAME}" "${DOWNLOAD_URL}"

# Make the binary executable (Linux and macOS only)
if [[ "$OS" == "linux"* ]] || [[ "$OS" == "darwin"* ]]; then
  chmod +x "${BINARY_NAME}"
fi

# Install the binary
echo "Installing ${BINARY_NAME} to ${INSTALL_DIR}..."
if [[ "$OS" == "msys"* ]] || [[ "$OS" == "cygwin"* ]] || [[ "$OS" == "mingw"* ]]; then
  mv "${BINARY_NAME}" "${INSTALL_DIR}"
else
  sudo mv "${BINARY_NAME}" "${INSTALL_DIR}"
fi

echo "Installation complete!"

