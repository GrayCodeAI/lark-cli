# Contributing to Lark CLI

Thank you for your interest in contributing to Lark CLI. This document covers the process for submitting changes.

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/<your-username>/lark-cli.git
   cd lark-cli
   ```
3. Create a branch:
   ```bash
   git checkout -b feature/my-change
   ```
4. Make your changes
5. Run tests:
   ```bash
   make test
   ```
6. Commit and push:
   ```bash
   git commit -m "feat: add my change"
   git push origin feature/my-change
   ```
7. Open a pull request against `main`

## Development

### Prerequisites

- Go 1.22 or later
- A running [Lark server](https://github.com/lark-dev/lark-core) for integration testing (optional)

### Building

```bash
make build          # Build binary to ./lark-cli
make install        # Install to $GOPATH/bin
make test           # Run tests with race detector
make lint           # Run go vet
make clean          # Remove build artifacts
```

### Code Style

- Follow standard Go conventions and `gofmt` formatting
- Keep command implementations in `internal/commands/`
- HTTP client logic belongs in `internal/client/`
- Use `text/tabwriter` for tabular CLI output
- Error messages should be lowercase and not end with punctuation

## Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add workflow trigger command
fix: handle 404 in list channels
docs: update installation instructions
refactor: extract shared tabwriter helper
test: add unit tests for client.Do
chore: update dependencies
```

## Pull Requests

- Keep PRs focused on a single change
- Include a clear description of what and why
- Ensure CI passes before requesting review
- Link related issues when applicable

## Reporting Bugs

Open an issue using the [bug report template](https://github.com/lark-dev/lark-cli/issues/new?template=bug_report.md). Include:

- Steps to reproduce
- Expected vs actual behavior
- CLI version (`lark-cli --version`)
- OS and architecture

## Feature Requests

Open an issue using the [feature request template](https://github.com/lark-dev/lark-cli/issues/new?template=feature_request.md). Describe the use case and proposed behavior.

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you agree to uphold its terms.
