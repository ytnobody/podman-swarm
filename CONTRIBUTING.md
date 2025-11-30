# Contributing to podman-swarm

Thank you for your interest in contributing to podman-swarm! This document provides guidelines for contributing to the project.

## Code of Conduct

Be respectful and constructive in all interactions with other contributors and maintainers.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git

### Setting Up Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/podman-swarm.git
   cd podman-swarm
   ```
3. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/ytnobody/podman-swarm.git
   ```

### Building and Testing

Build the project:
```bash
go build -o podman-swarm
```

Run tests:
```bash
go test ./...
```

Run with verbose output:
```bash
go test -v ./...
```

## Development Workflow

1. Create a new branch for your feature or fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the code style guidelines

3. Add or update tests as needed

4. Ensure all tests pass:
   ```bash
   go test ./...
   ```

5. Commit your changes with clear commit messages:
   ```bash
   git commit -m "Brief description of changes"
   ```

6. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

7. Open a Pull Request on GitHub with a clear description of your changes

## Code Style Guidelines

### Go Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code:
  ```bash
  gofmt -w .
  ```
- Run `go vet` to catch common errors:
  ```bash
  go vet ./...
  ```

### Naming Conventions

- Use clear, descriptive names for variables, functions, and packages
- Follow Go naming conventions (CamelCase for exported identifiers, lowercase for unexported)
- Avoid single-letter variable names except for loops and mathematical contexts

### Comments

- Add comments for exported functions and types
- Keep comments concise and clear
- Update comments when code changes

## Testing Requirements

- Write tests for new features and bug fixes
- Ensure all tests pass before submitting a Pull Request
- Add unit tests in `cmd/commands_test.go` or appropriate `*_test.go` files
- Use the test utilities in `cmd/internal/test/` for mock objects

## Pull Request Process

1. Update README.md or relevant documentation if needed
2. Ensure your PR title clearly describes the changes
3. Provide a clear description of what your PR does
4. Reference any related issues (e.g., "Fixes #123")
5. Be responsive to code review feedback

## Reporting Issues

### Bug Reports

When reporting a bug, please include:
- Go version (`go version`)
- Operating system
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Any error messages or logs

### Feature Requests

For feature requests, please:
- Describe the use case
- Explain the expected behavior
- Provide any relevant context

## License

By contributing to podman-swarm, you agree that your contributions will be licensed under its MIT License.

## Questions?

Feel free to open a discussion or issue on GitHub if you have questions about contributing.

Thank you for contributing to podman-swarm!
