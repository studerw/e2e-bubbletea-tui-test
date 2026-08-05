# TODO

## Airgapped Environment Support
- [ ] Add `API_BASE_URL` environment variable support to the webcall panel so the base URL can be overridden (default: `https://jsonplaceholder.typicode.com`)
- [ ] Create a local mock HTTP server (Go or Node.js) that returns canned JSON responses matching the JSONPlaceholder API
- [ ] Update Shell-Use E2E tests to start the mock server in `beforeAll` and set `API_BASE_URL` when launching the TUI binary
- [ ] Update Go-level `teatest` tests to use the mock server or inject test data

## CI/CD
- [ ] Add GitLab CI configuration (`.gitlab-ci.yml`) for running both E2E test suites
- [ ] Ensure `shell-use` CLI is installed in the CI runner image
- [ ] Ensure Node.js 20+ is available in the CI runner image
- [ ] Add CI caching for `node_modules` and Go module cache

## Future Test Improvements
- [ ] Add screenshot capture (`shell-use screenshot`) at key points for visual regression artifacts
- [ ] Add Shell-Use session recording for debugging test failures
- [ ] Consider adding more granular todo list tests (toggle completion, delete items)
- [ ] Consider adding timer lap recording tests
- [ ] Add Web Call loading state assertion (spinner visible during fetch)

## EnterAltScreen()
- Need to figure out why this was removed from e2e tests and why we still need it (ask Ryan)
