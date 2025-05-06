#!/bin/bash

# バージョン
VERSION="1.0.0"

# OSとアーキテクチャの検出
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
  ARCH="arm64"
fi

# ダウンロードURL
DOWNLOAD_URL="https://github.com/yk68-hanpen/passgen/releases/download/v${VERSION}/passgen-${OS}-${ARCH}"

# インストール先
INSTALL_DIR="/usr/local/bin"
INSTALL_PATH="${INSTALL_DIR}/passgen"

# インストール先ディレクトリの存在確認と作成
if [ ! -d "$INSTALL_DIR" ]; then
  echo "Creating directory $INSTALL_DIR..."
  mkdir -p "$INSTALL_DIR" 2>/dev/null || sudo mkdir -p "$INSTALL_DIR"
fi

# インストール
echo "Downloading passgen v${VERSION} for ${OS}-${ARCH}..."
curl -L -s "${DOWNLOAD_URL}" -o /tmp/passgen || { echo "Download failed"; exit 1; }

echo "Installing to ${INSTALL_PATH}..."
mv /tmp/passgen "${INSTALL_PATH}" 2>/dev/null || sudo mv /tmp/passgen "${INSTALL_PATH}" || { echo "Installation failed"; exit 1; }

chmod +x "${INSTALL_PATH}" 2>/dev/null || sudo chmod +x "${INSTALL_PATH}" || { echo "Setting executable permission failed"; exit 1; }

echo "✅ passgen has been installed to ${INSTALL_PATH}"
echo "Run 'passgen --help' to get started"
