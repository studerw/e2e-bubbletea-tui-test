# SPEC

We are building a MVP for e2e testing of a TUI application. Our goal is to create a small TUI demo app and then create some E2E tests that will 
run the app and use some library e.g. [Shell-Use](https://github.com/microsoft/shell-use) and Node bindings to test interacting with the TUI app.

This doc should be updated as we add to the application.

## Requirements

* Written in Go (1.26.5)
* Uses [BubbleTea v2 library](https://github.com/charmbracelet/bubbletea) — stable release via `charm.land/bubbletea/v2`
* Uses [Lipgloss v2](https://github.com/charmbracelet/lipgloss) — stable release via `charm.land/lipgloss/v2`
* App uses the Alt Screen mode (declarative via `tea.View`)
* A makefile is created to run the various necessary goals:
    * build the app
    * run the app
    * lint the app
    * run any unit tests
    * run the full e2e tests (hybrid: Go `teatest` + Shell-Use Node.js)

## E2E Testing

A **hybrid** approach is used for end-to-end testing:

* **Go-level tests (`teatest`)**: Fast, programmatic tests using BubbleTea's `teatest` package from `charmbracelet/x/exp/teatest/v2`. Tests tab navigation, panel content rendering, user interactions (typing, selecting, toggling), and quit behavior without a real terminal.
* **Shell-Use tests (Node.js/Vitest)**: Integration tests using [Shell-Use](https://github.com/microsoft/shell-use) CLI + `@microsoft/shell-use` Node.js bindings. Launches the real compiled binary in a real terminal session, sends keystrokes, and asserts on actual rendered terminal text output.
* Tests run sequentially: `make e2e-go` then `make e2e-shell`, or combined via `make e2e`
* The Web Call panel currently uses the public JSONPlaceholder API. See [TODO.md](TODO.md) for airgapped environment support.

## App

* The main app is a standard Go app that runs from a standard shell and opens up a TUI that is compatible in any standard Mac, Linux, and Windows Terminal
    * On a similar project, we ended up having the wrap the e2e tests in [Cmder](https://cmder.app/) but unless this is necessary, we don't want to do this unless absolutely necessary.

## TUI

* The main TUI will consist of a:
    * header e.g. _E2E MVC using BubbleTea v2_ that's about 10% of the top screen, full width
    * a main middle panel, full width, that displays the current panel that's about 80% of the screen
    * a bottom footer, full width, with some dummy text e.g. Copyright Clarity Innovations or something like most footers have with appropriate unicode TM / Copyright characters

* Main initial panel contains tabs that are keyboard shortcut accessible e.g. tab and enter to choose. We will start with 4 tabs and as we add more, this doc should be updated. All tab panels, when opened, should show the tab visually that is choosen and the panel changes when they click a different tab
    * Tab 1: `Text` - opens a panel that has some lorem ipsum text that we can test that it appears when clicking the `Text` tab button from the main tabs panel
    * Tab 2: `Web Call` - make a webcall with a select of options to choose using some web testing site that grabs some deterministic data based on the option choosen and then displays the response. We might need to discuss which site we can use for this test. The test of this tab will be that the user clicks a certain option in the drop-down selector and then waits for the call the respond and asserts that the correct payload or data is now showing up in the tab's panel.
    * Tab 3: Give me some options using [BubbleTea Examples](https://github.com/charmbracelet/bubbletea/tree/main/examples) that we can use for our test. It should be a thorough end to end test that is more complex than just displaying text and asserting it's there on the screen
    * Tab 4: Give me some options, same as step 3.

## Code Style
* I am a begginer Go coder, so I want you use the most standard Go project and coding styles, formatting, etc. Include typical linter info and other supporting buildfiles and anything else a typical Go project would use, and tie them to the Makefile. 

## README
There should be a standard README.md that gives a concise summary of the project and how to use the Makefile and any other pertinent information for developers and testers. 




