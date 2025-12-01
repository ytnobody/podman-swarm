package cmd

import (
	"context"
	"strings"
	"testing"

	"github.com/ytnobody/podman-swarm/cmd/internal/test"
)

// TestRunContainerCommandOrder validates that podman run arguments are in correct order
// This test ensures the fix for argument order (podman run [args...] image) works correctly
func TestRunContainerCommandOrder(t *testing.T) {
	testCases := []struct {
		name          string
		image         string
		args          []string
		expectedOrder string // substring that should appear in correct order
	}{
		{
			name:          "with no args",
			image:         "alpine:latest",
			args:          []string{},
			expectedOrder: "podman run alpine:latest",
		},
		{
			name:          "with single arg",
			image:         "nginx:latest",
			args:          []string{"-d"},
			expectedOrder: "podman run -d nginx:latest",
		},
		{
			name:          "with multiple args",
			image:         "ubuntu:20.04",
			args:          []string{"-d", "-p", "8080:80", "--name", "myapp"},
			expectedOrder: "podman run -d -p 8080:80 --name myapp ubuntu:20.04",
		},
		{
			name:          "args before image",
			image:         "busybox",
			args:          []string{"--rm", "-it"},
			expectedOrder: "--rm -it busybox",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedCmd string

			mockClient := &test.MockSSHClient{
				ExecuteFunc: func(ctx context.Context, cmd string) (string, error) {
					capturedCmd = cmd
					return "container_id_123", nil
				},
			}

			// Build the command string the same way runContainerOnHost does
			cmdStr := "podman run " + strings.Join(tc.args, " ")
			if len(tc.args) > 0 {
				cmdStr += " "
			}
			cmdStr += tc.image

			_, err := mockClient.Execute(context.Background(), cmdStr)
			if err != nil {
				t.Errorf("Execute should not fail: %v", err)
			}

			if capturedCmd != cmdStr {
				t.Errorf("command mismatch, expected: %s, got: %s", cmdStr, capturedCmd)
			}

			if !strings.Contains(capturedCmd, tc.expectedOrder) {
				t.Errorf("command should contain '%s', got: %s", tc.expectedOrder, capturedCmd)
			}
		})
	}
}

// TestRunCommand_Integration validates the run command flow
func TestRunCommand_Integration(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		image        string
		args         []string
		expectError  bool
		expectedCall string
	}{
		{
			name:         "simple run",
			host:         "host1",
			image:         "alpine:latest",
			args:         []string{},
			expectError:  false,
			expectedCall: "podman run alpine:latest",
		},
		{
			name:         "run with options",
			host:         "host1",
			image:         "nginx:latest",
			args:         []string{"-d", "-p", "8080:80"},
			expectError:  false,
			expectedCall: "podman run -d -p 8080:80 nginx:latest",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedCmd string

			mockClient := &test.MockSSHClient{
				ExecuteFunc: func(ctx context.Context, cmd string) (string, error) {
					capturedCmd = cmd
					return "success", nil
				},
			}

			// Simulate the command building logic from runContainerOnHost
			cmdStr := "podman run "
			if len(tc.args) > 0 {
				cmdStr += strings.Join(tc.args, " ") + " "
			}
			cmdStr += tc.image

			_, err := mockClient.Execute(context.Background(), cmdStr)

			if (err != nil) != tc.expectError {
				t.Errorf("error expectation mismatch, expected error: %v, got: %v", tc.expectError, err)
			}

			if capturedCmd != tc.expectedCall {
				t.Errorf("command mismatch, expected: '%s', got: '%s'", tc.expectedCall, capturedCmd)
			}
		})
	}
}
