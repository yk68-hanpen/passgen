# PassGen

PassGen is a simple command-line tool for generating strong passwords.

## Features

- Generate one or multiple passwords at once
- Customize password length
- Include or exclude character types (lowercase, uppercase, numbers, special characters)
- Simple and easy-to-use CLI interface

## Installation

### Easy Install (macOS/Linux)

```bash
curl -L https://raw.githubusercontent.com/yk68-hanpen/passgen/main/install.sh | bash
```

### Manual Installation

#### Download Binary

Download the appropriate binary for your platform from the [Releases](https://github.com/yk68-hanpen/passgen/releases) page.

#### Build from Source

```bash
git clone https://github.com/yk68-hanpen/passgen.git
cd passgen
go build -o passgen
sudo mv passgen /usr/local/bin/
```

## Usage

```bash
# Generate a single password with default options
passgen gen

# Generate multiple passwords
passgen gen -c 5

# Generate a password with custom length
passgen gen -l 16

# Generate a password without special characters
passgen gen -s=false

# Generate multiple passwords with custom options
passgen gen -c 3 -l 20 -s=false -n=true
```

### Available Options

- `-c, --count`: Number of passwords to generate (default: 1)
- `-l, --length`: Length of the password (default: 12)
- `-s, --special`: Include special characters (default: true)
- `-n, --numbers`: Include numbers (default: true)
- `-u, --uppercase`: Include uppercase letters (default: true)
- `-w, --lowercase`: Include lowercase letters (default: true)

## License

MIT
