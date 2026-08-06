# TUI E2E Demo

A TUI application built with Go and BubbleTea v2 for end-to-end testing purposes.

I used _Go_ version _1.26.5_ and the corresponding _golangci-lint_. 

## Prerequisites

- **Go** 1.26 ([install](https://go.dev/dl/))
- **golangci-lint** (optional, for linting): `make install-tools`
  
For Shell-Use E2E tests only:
- **Node.js** 20+ ([install](https://nodejs.org/))
- **shell-use** CLI — a Rust-powered terminal automation tool from [microsoft/shell-use](https://github.com/microsoft/shell-use). The Node.js `@microsoft/shell-use` package is a client that talks to this CLI binary, so it **must** be installed separately:

  **macOS (Homebrew):**
  ```bash
  brew tap microsoft/shell-use https://github.com/microsoft/shell-use
  brew install shell-use
  ```

  **macOS / Linux (install script):**
  ```bash
  curl --proto '=https' --tlsv1.2 -LsSf https://raw.githubusercontent.com/microsoft/shell-use/main/install/install.sh | sh
  ```

  **Windows (PowerShell):**
  ```powershell
  irm https://raw.githubusercontent.com/microsoft/shell-use/main/install/install.ps1 | iex
  ```

  **Windows (winget):**
  ```bash
  winget install Microsoft.ShellUse
  ```

  Verify it's installed: `shell-use --version`

## Quick Start

```bash
make deps    # Download Go dependencies
make run     # Build and run the TUI app
```

## Local Setup (First Time)

```bash
# 1. Clone and install Go dependencies
make deps

# 2. Build and verify the app runs
make build
make run          # Press q to quit

# 3. Run Go-level E2E tests (no extra dependencies needed)
make e2e-go

# 4. (Optional) Set up Shell-Use tests
#    Requires: Node.js 20+ and shell-use CLI (see Prerequisites)
make e2e-install  # Install Node.js test dependencies
make e2e-shell    # Run Shell-Use E2E tests

# 5. Run everything
make e2e          # Runs both Go and Shell-Use E2E tests
```

## Testing

This project uses a **hybrid E2E testing approach**:

### Go-Level Tests (`teatest`)
Fast, programmatic tests that drive the BubbleTea model directly without a real terminal. Tests tab navigation, panel content, user interactions, and quit behavior.

```bash
make e2e-go
```

### Shell-Use Tests (Node.js + Vitest)
Integration tests that launch the real compiled binary in a terminal session using [Shell-Use](https://github.com/microsoft/shell-use). Tests the actual rendered TUI output with real keystrokes. Requires the `shell-use` CLI binary and Node.js 20+.

```bash
make e2e-install  # First time only: install Node.js dependencies
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

## Shell-Use Debugging

Shell-Use runs in a **headless PTY** — there's no visible terminal window during tests. Here's how to observe what's happening:

### Read the terminal buffer

Use Shell-Use's `text()` method to dump the current screen contents from inside any test:

```typescript
const screenText = await su.text();
console.log('--- SCREEN ---');
console.log(screenText);
console.log('--- END ---');
```

Output appears in the Vitest console when using `--reporter=verbose` (already configured).

### Dump screen on failure

Add an `afterEach` hook that captures the terminal buffer only when a test fails:

```typescript
import { afterEach } from 'vitest';

afterEach(async (context) => {
  if (context.task.result?.state === 'fail') {
    try {
      const screenText = await su.text();
      console.log(`\n=== SCREEN ON FAILURE (${context.task.name}) ===`);
      console.log(screenText);
      console.log('=== END ===\n');
    } catch {
      // session may already be closed
    }
  }
});
```

### Stop on first failure

Use `--bail 1` to stop the test run immediately on the first failure so you can inspect the output:

```bash
cd tests/e2e-shell-use && npx vitest run --reporter=verbose --bail 1
```

### Running Specific Tests

You can target specific files or test cases rather than running the whole suite:

**Run a single test file:**
```bash
cd tests/e2e-shell-use && npx vitest run src/app-layout.test.ts
```

**Run a single test case (by name):**
```bash
cd tests/e2e-shell-use && npx vitest run -t "displays the header"
```

**Run a specific test case inside a specific file:**
```bash
cd tests/e2e-shell-use && npx vitest run src/app-layout.test.ts -t "displays the header"
```

### Manual side-by-side testing

For visual debugging, open two terminal windows:
1. **Window 1**: Run the TUI manually — `make run`
2. **Window 2**: Step through the same keystrokes the tests send, observing the real rendered output

This is the simplest way to visually verify what the tests are asserting against.

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
