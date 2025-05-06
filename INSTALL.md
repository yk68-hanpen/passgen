# Installation Instructions

## Easy Install (macOS/Linux)

The easiest way to install PassGen is using our installation script:

```bash
curl -L https://raw.githubusercontent.com/yk68-hanpen/passgen/main/install.sh | bash
```

This script will:
1. Detect your operating system and architecture
2. Download the appropriate binary from GitHub Releases
3. Install it to `/usr/local/bin/passgen`
4. Make it executable

## Manual Installation

### Prerequisites

- Go 1.21 or later (only if building from source)

### Option 1: Download Binary

1. Go to the [Releases](https://github.com/yk68-hanpen/passgen/releases) page
2. Download the appropriate binary for your platform:
   - macOS Intel: `passgen-darwin-amd64`
   - macOS Apple Silicon: `passgen-darwin-arm64`
   - Linux x86_64: `passgen-linux-amd64`
   - Linux ARM64: `passgen-linux-arm64`

3. Make it executable and move it to your PATH:
   ```bash
   chmod +x passgen-*
   sudo mv passgen-* /usr/local/bin/passgen
   ```

### Option 2: Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yk68-hanpen/passgen.git
   cd passgen
   ```

2. Build the binary:
   ```bash
   go build -o passgen
   ```

3. Move the binary to a directory in your PATH:
   ```bash
   sudo mv passgen /usr/local/bin/
   ```

## Verifying Installation

To verify that PassGen is installed correctly, run:

```bash
passgen --help
```

You should see the help output describing the available commands and options.
