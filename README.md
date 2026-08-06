# TUI E2E Demo

This project is a dedicated sandbox designed to explore and evaluate different End-to-End (E2E) testing frameworks for Terminal User Interfaces (TUIs). It uses [BubbleTea v2](https://github.com/charmbracelet/bubbletea) for interactive UI construction and [Lipgloss v2](https://github.com/charmbracelet/lipgloss) for TUI styling and layouts.

## Project Goal
The primary objective of this repository is to demonstrate how to test terminal applications under two distinct testing paradigms:
1. **In-Process Testing**: Driving the Bubble Tea model programmatically without a real terminal.
2. **Terminal Automation Testing**: Driving the compiled binary in an automated terminal emulator (headless PTY).

## Quick Links to Test Suites
- **Go E2E Tests ([teatest](https://pkg.go.dev/github.com/charmbracelet/x/exp/teatest/v2))**: Located in [`tests/e2e/app_test.go`](./tests/e2e/app_test.go) — programmatically sends events and asserts state in-process.
- **Shell-Use E2E Tests (Node.js/Vitest via [Shell-Use](https://github.com/microsoft/shell-use))**: Located in the [`tests/e2e-shell-use/`](./tests/e2e-shell-use/) directory — drives the compiled terminal app inside a PTY using `@microsoft/shell-use`.

---

## Docker Quick Start

> **Only Docker is required** — none of the build and run tools need to be installed locally.

```bash
# 1. Build the Docker dev image (one-time; re-run when Dockerfile/go.mod/package-lock.json change)
make build-docker-image

# 2. Compile the app binary inside Docker (writes to bin/ — same as 'make build' locally)
make docker-build

# 3. Run the TUI interactively
make docker-run

# 4. Run all tests (unit + E2E) — fully headless, no TTY needed
make docker-test
make docker-e2e

# 5. Debug: open a shell inside the container
make docker-shell
```

See the [Docker Commands](#docker-commands) table below for the full list of targets.

## Prerequisites

### Using Docker (recommended)

Install **[Docker](https://docs.docker.com/get-docker/)** — that's it. Run `make docker-build` once to create the dev image, then use `make docker-*` targets for everything else.

### Running natively (optional)

If you prefer to run without Docker:

- **Go** 1.26 ([install](https://go.dev/dl/))
- **golangci-lint** (optional, for linting): `make install-tools`
  
For Shell-Use E2E tests only:
- **Node.js** 22+ ([install](https://nodejs.org/))
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

### Docker Commands

Only Docker is required for these targets — no Go, Node, or other tools needed.

| Command | Description |
|---------|-------------|
| `make build-docker-image` | Build the dev image (`tui-e2e-demo:dev`) — one-time setup |
| `make docker-build` | Compile the app binary inside Docker → writes to `bin/` (like `make build` locally). Note that on e.g. Mac, the compiled binary wouldn't be runnable on your native machine - only in the image itself. |
| `make docker-run` | Run the TUI interactively inside Docker (**requires a TTY; local use only**) |
| `make docker-test` | Run unit tests inside Docker (headless / CI-compatible) |
| `make docker-lint` | Run golangci-lint inside Docker (headless / CI-compatible) |
| `make docker-e2e` | Run all E2E tests inside Docker (headless / CI-compatible) |
| `make docker-e2e-go` | Run Go/teatest E2E tests inside Docker |
| `make docker-e2e-shell` | Run Shell-Use/Vitest E2E tests inside Docker |
| `make docker-shell` | Open a bash shell inside the container for debugging |

> **`DOCKER_IMAGE` override**: set the variable to push/pull from a registry:
> ```bash
> make build-docker-image DOCKER_IMAGE=registry.example.com/myteam/tui-e2e-demo:latest
> ```

## Running in CI (GitLab CI / GitHub Actions)

All `docker-e2e*` and `docker-test` / `docker-lint` targets are **fully headless** — they do not require an allocated TTY, so they work unmodified in any Docker-capable CI runner.

### Why headless works

| Test suite | How it avoids needing a host TTY |
|---|---|
| Go unit tests | Pure in-process, no terminal needed |
| Go E2E (`teatest`) | `teatest` creates an **in-process** pseudo-terminal and drives BubbleTea programmatically; no host TTY is allocated |
| Shell-Use / Vitest | `shell-use` spawns its own PTY via `/dev/pts`, which is mounted inside every Docker container by default; the test process itself runs headlessly |

### Example: GitLab CI (`.gitlab-ci.yml`)

```yaml
stages:
  - build
  - test

variables:
  DOCKER_IMAGE: $CI_REGISTRY_IMAGE:$CI_COMMIT_SHORT_SHA

build-image:
  stage: build
  image: docker:26
  services:
    - docker:26-dind
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - make build-docker-image DOCKER_IMAGE=$DOCKER_IMAGE
    - docker push $DOCKER_IMAGE

unit-tests:
  stage: test
  image: $DOCKER_IMAGE
  script:
    - make test    # runs natively inside the already-correct image

e2e-tests:
  stage: test
  image: docker:26
  services:
    - docker:26-dind
  script:
    - docker pull $DOCKER_IMAGE
    - make docker-e2e DOCKER_IMAGE=$DOCKER_IMAGE
```

> **Simpler alternative**: if the CI runner itself uses your dev image, just call `make test`, `make e2e-go`, and `make e2e-shell` directly — no extra Docker-in-Docker setup needed.

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


