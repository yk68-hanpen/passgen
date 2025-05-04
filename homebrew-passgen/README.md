# Homebrew Tap for PassGen

This repository contains the Homebrew formula for [PassGen](https://github.com/yk68-hanpen/passgen), a simple password generator CLI tool.

## Installation

```bash
# Add the tap repository
brew tap yk68-hanpen/passgen

# Install PassGen
brew install passgen
```

## Usage

Once installed, you can use PassGen to generate passwords:

```bash
# Generate a single password with default options
passgen gen

# Generate multiple passwords
passgen gen -c 5

# Generate a password with custom length
passgen gen -l 16
```

For more information, see the [PassGen repository](https://github.com/yk68-hanpen/passgen).
