# TUI E2E Demo

A TUI application built with Go and BubbleTea v2 for end-to-end testing purposes.

I used _Go_ version _1.26.5_ and the corresponding _golangci-lint_. 

## Quick Start

```bash
make deps    # Download dependencies
make run     # Build and run
```

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make all` | Run lint, test, and build |
| `make build` | Build the application |
| `make run` | Build and run |
| `make test` | Run unit tests |
| `make lint` | Run linter |
| `make e2e` | Run E2E tests |
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

© 2024 Clarity Innovations™
