# TUI E2E Demo

A TUI application built with Go and BubbleTea v2 for end-to-end testing purposes.

I used _Go_ version _1.26.5_ and the corresponding _golangci-lint_. 

## Prerequisites

- **Go** 1.22+
- **Node.js** 20+ (for Shell-Use E2E tests)
- **shell-use** CLI ([microsoft/shell-use](https://github.com/microsoft/shell-use)) — install via:
  ```bash
  brew tap microsoft/shell-use https://github.com/microsoft/shell-use
  brew install shell-use
  ```
- **golangci-lint** (optional, for linting): `make install-tools`

## Quick Start

```bash
make deps    # Download Go dependencies
make run     # Build and run the TUI app
```

## Testing

This project uses a **hybrid E2E testing approach**:

### Go-Level Tests (`teatest`)
Fast, programmatic tests that drive the BubbleTea model directly without a real terminal. Tests tab navigation, panel content, user interactions, and quit behavior.

```bash
make e2e-go
```

### Shell-Use Tests (Node.js + Vitest)
Integration tests that launch the real compiled binary in a terminal session using [Shell-Use](https://github.com/microsoft/shell-use). Tests the actual rendered TUI output with real keystrokes.

```bash
make e2e-install  # First time: install Node.js dependencies
make e2e-shell    # Run Shell-Use E2E tests
```

### Run All E2E Tests
```bash
make e2e          # Runs both Go and Shell-Use E2E tests
```

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make all` | Run lint, test, and build |
| `make build` | Build the application |
| `make run` | Build and run |
| `make test` | Run unit tests |
| `make lint` | Run linter |
| `make e2e` | Run all E2E tests (Go + Shell-Use) |
| `make e2e-go` | Run Go-level E2E tests (teatest) |
| `make e2e-shell` | Run Shell-Use E2E tests (Node.js) |
| `make e2e-install` | Install Shell-Use test dependencies |
| `make clean` | Remove build artifacts |
| `make deps` | Download dependencies |
| `make help` | Show all commands |

## Keyboard Shortcuts

- `q` / `Ctrl+C`: Quit
- `Tab` / `Shift+Tab`: Navigate tabs
- `1-4`: Jump to tab

## Tabs

1. **Text**: Lorem Ipsum display
2. **Web Call**: JSONPlaceholder API requests
3. **Todo List**: Interactive task manager
4. **Timer**: Stopwatch with laps

© 2024 Clarity Innovations
