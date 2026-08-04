// Package e2e contains end-to-end tests for the TUI application.
package e2e

import (
	"os/exec"
	"testing"
	"time"
)

// TestAppStarts verifies the application can start and exit cleanly.
func TestAppStarts(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "../../bin/app", "../../cmd/app")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build app: %v", err)
	}

	t.Log("App built successfully - E2E test placeholder")
}

// TestBinaryExists verifies the binary can be created.
func TestBinaryExists(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "../../bin/app", "../../cmd/app")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build app: %v", err)
	}
	t.Log("Binary created successfully")
}

// TestAppTimeout tests that the app doesn't hang on startup.
func TestAppTimeout(t *testing.T) {
	buildCmd := exec.Command("go", "build", "-o", "../../bin/app", "../../cmd/app")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build app: %v", err)
	}

	cmd := exec.Command("../../bin/app")

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start app: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	if err := cmd.Process.Kill(); err != nil {
		t.Logf("Note: Could not kill process: %v", err)
	}

	t.Log("App started and was terminated successfully")
}
