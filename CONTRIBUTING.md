# Contributing to PassGen

Thank you for your interest in contributing to PassGen! Here are some guidelines to help you get started.

## Development Setup

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/yk68-hanpen/passgen.git
   cd passgen
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Making Changes

1. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Run tests:
   ```bash
   go test ./...
   ```
4. Commit your changes:
   ```bash
   git commit -m "Add your meaningful commit message"
   ```
5. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
6. Create a Pull Request

## Code Style

- Follow standard Go code style and conventions
- Use `gofmt` to format your code
- Add comments for public functions and packages

## Testing

- Add tests for new features
- Ensure all tests pass before submitting a PR

## Releasing

Releases are managed by the maintainers. If you think a new release is needed, please open an issue.

## License

By contributing to PassGen, you agree that your contributions will be licensed under the project's MIT License.
