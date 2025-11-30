package cmd

import (
	"context"
	"errors"
	"testing"
	
	"github.com/ytnobody/podman-swarm/cmd/internal/test"
)

// Test inspect command argument validation
func TestInspectCommand_ValidArgs(t *testing.T) {
	// Should not return arg validation error
	if err := inspectCmd.Args(inspectCmd, []string{"host1", "container1"}); err != nil {
		t.Errorf("inspect command with valid args should not error, got: %v", err)
	}
}

func TestInspectCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := inspectCmd.Args(inspectCmd, []string{}); err == nil {
		t.Error("inspect command with no args should return an error")
	}
	
	// Test with only 1 arg
	if err := inspectCmd.Args(inspectCmd, []string{"host1"}); err == nil {
		t.Error("inspect command with only 1 arg should return an error")
	}
}

// Test run command argument validation
func TestRunCommand_ValidArgs(t *testing.T) {
	// Valid with host and image
	if err := runCmd.Args(runCmd, []string{"host1", "alpine:latest"}); err != nil {
		t.Errorf("run command with host and image should succeed, got: %v", err)
	}
	
	// Valid with host, image and options
	if err := runCmd.Args(runCmd, []string{"host1", "alpine:latest", "-d"}); err != nil {
		t.Errorf("run command with host, image and options should succeed, got: %v", err)
	}
}

func TestRunCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := runCmd.Args(runCmd, []string{}); err == nil {
		t.Error("run command with no args should return an error")
	}
	
	// Test with only host
	if err := runCmd.Args(runCmd, []string{"host1"}); err == nil {
		t.Error("run command with only host should return an error")
	}
}

// Test stop command argument validation
func TestStopCommand_ValidArgs(t *testing.T) {
	if err := stopCmd.Args(stopCmd, []string{"host1", "container1"}); err != nil {
		t.Errorf("stop command with host and container should succeed, got: %v", err)
	}
}

func TestStopCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := stopCmd.Args(stopCmd, []string{}); err == nil {
		t.Error("stop command with no args should return an error")
	}
	
	// Test with only host
	if err := stopCmd.Args(stopCmd, []string{"host1"}); err == nil {
		t.Error("stop command with only host should return an error")
	}
}

// Test rm command argument validation
func TestRmCommand_ValidArgs(t *testing.T) {
	if err := rmCmd.Args(rmCmd, []string{"host1", "container1"}); err != nil {
		t.Errorf("rm command with host and container should succeed, got: %v", err)
	}
}

func TestRmCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := rmCmd.Args(rmCmd, []string{}); err == nil {
		t.Error("rm command with no args should return an error")
	}
	
	// Test with only host
	if err := rmCmd.Args(rmCmd, []string{"host1"}); err == nil {
		t.Error("rm command with only host should return an error")
	}
}

// Test exec command argument validation
func TestExecCommand_ValidArgs(t *testing.T) {
	if err := execCmd.Args(execCmd, []string{"host1", "container1", "ls"}); err != nil {
		t.Errorf("exec command with host, container and command should succeed, got: %v", err)
	}
}

func TestExecCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := execCmd.Args(execCmd, []string{}); err == nil {
		t.Error("exec command with no args should return an error")
	}
	
	// Test with only host
	if err := execCmd.Args(execCmd, []string{"host1"}); err == nil {
		t.Error("exec command with only host should return an error")
	}
	
	// Test with only host and container
	if err := execCmd.Args(execCmd, []string{"host1", "container1"}); err == nil {
		t.Error("exec command with only host and container should return an error")
	}
}

// Test MockSSHClient
func TestMockSSHClient_Execute(t *testing.T) {
	expectedOutput := "test output"
	
	client := &test.MockSSHClient{
		ExecuteFunc: func(ctx context.Context, cmd string) (string, error) {
			return expectedOutput, nil
		},
	}
	
	output, err := client.Execute(context.Background(), "test command")
	if err != nil {
		t.Errorf("execute should not error, got: %v", err)
	}
	if output != expectedOutput {
		t.Errorf("execute output mismatch, expected: %s, got: %s", expectedOutput, output)
	}
}

func TestMockSSHClient_ExecuteError(t *testing.T) {
	expectedErr := errors.New("connection failed")
	
	client := &test.MockSSHClient{
		ExecuteFunc: func(ctx context.Context, cmd string) (string, error) {
			return "", expectedErr
		},
	}
	
	_, err := client.Execute(context.Background(), "test command")
	if err != expectedErr {
		t.Errorf("execute error mismatch, expected: %v, got: %v", expectedErr, err)
	}
}

// Test logs command argument validation
func TestLogsCommand_ValidArgs(t *testing.T) {
	if err := logsCmd.Args(logsCmd, []string{"host1", "container1"}); err != nil {
		t.Errorf("logs command with host and container should succeed, got: %v", err)
	}
}

func TestLogsCommand_MissingArgs(t *testing.T) {
	// Test with no args
	if err := logsCmd.Args(logsCmd, []string{}); err == nil {
		t.Error("logs command with no args should return an error")
	}
	
	// Test with only host
	if err := logsCmd.Args(logsCmd, []string{"host1"}); err == nil {
		t.Error("logs command with only host should return an error")
	}
}

// Test MockConfig
func TestMockConfig_GetHostByName(t *testing.T) {
	cfg := test.MockConfig()
	
	host := cfg.GetHostByName("host1")
	if host == nil {
		t.Error("GetHostByName should return host1")
	}
	if host.Address != "192.168.1.10" {
		t.Errorf("host1 address mismatch, expected: 192.168.1.10, got: %s", host.Address)
	}
}

func TestMockConfig_GetHostByNameNotFound(t *testing.T) {
	cfg := test.MockConfig()
	
	host := cfg.GetHostByName("nonexistent")
	if host != nil {
		t.Error("GetHostByName should return nil for nonexistent host")
	}
}

func TestMockConfig_GetHostOrGroup(t *testing.T) {
	cfg := test.MockConfig()
	
	// Test single host
	hosts := cfg.GetHostOrGroup("host1")
	if len(hosts) != 1 {
		t.Errorf("GetHostOrGroup should return 1 host, got: %d", len(hosts))
	}
	
	// Test group
	hosts = cfg.GetHostOrGroup("all")
	if len(hosts) != 2 {
		t.Errorf("GetHostOrGroup should return 2 hosts for 'all' group, got: %d", len(hosts))
	}
}
