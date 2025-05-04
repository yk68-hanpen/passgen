# Installation Instructions

## Using Homebrew (macOS)

The easiest way to install PassGen on macOS is using Homebrew:

```bash
# Add the tap repository
brew tap yk68-hanpen/passgen

# Install PassGen
brew install passgen
```

## Manual Installation

### Prerequisites

- Go 1.21 or later

### Building from Source

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
