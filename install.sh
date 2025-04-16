#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

# Version to install
VERSION="v0.1.1"
BINARY_NAME="bucketx"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Convert architecture names
case ${ARCH} in
    x86_64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture: ${ARCH}${NC}"
        exit 1
        ;;
esac

# Set archive extension based on OS
if [ "$OS" = "windows" ]; then
    ARCHIVE_EXT="zip"
else
    ARCHIVE_EXT="tar.gz"
fi

# Construct download URL
DOWNLOAD_URL="https://github.com/TeamXSeven/bucketX/releases/download/${VERSION}/bucketx_${OS}_${ARCH}.${ARCHIVE_EXT}"

echo -e "${BLUE}Installing BucketX ${VERSION}...${NC}"
echo -e "${BLUE}Downloading from: ${DOWNLOAD_URL}${NC}"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download the archive
if command -v curl > /dev/null; then
    curl -L -o "${BINARY_NAME}.${ARCHIVE_EXT}" "${DOWNLOAD_URL}"
else
    wget -O "${BINARY_NAME}.${ARCHIVE_EXT}" "${DOWNLOAD_URL}"
fi

# Extract the archive
if [ "$ARCHIVE_EXT" = "zip" ]; then
    unzip "${BINARY_NAME}.${ARCHIVE_EXT}"
else
    tar xzf "${BINARY_NAME}.${ARCHIVE_EXT}"
fi

# Install the binary
sudo mv "${BINARY_NAME}" /usr/local/bin/
sudo chmod +x /usr/local/bin/"${BINARY_NAME}"

# Clean up
cd - > /dev/null
rm -rf "$TMP_DIR"

echo -e "${GREEN}BucketX has been installed successfully!${NC}"
echo -e "${BLUE}You can now use it by running: ${NC}bucketx"