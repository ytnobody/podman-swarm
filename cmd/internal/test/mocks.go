package test

import (
	"context"
	"github.com/ytnobody/podman-swarm/pkg/config"
	"github.com/ytnobody/podman-swarm/pkg/ssh"
)

// MockSSHClient is a mock implementation of ssh.Client
type MockSSHClient struct {
	ExecuteFunc func(ctx context.Context, cmd string) (string, error)
	CloseFunc   func() error
}

func (m *MockSSHClient) Execute(ctx context.Context, cmd string) (string, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(ctx, cmd)
	}
	return "", nil
}

func (m *MockSSHClient) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}

// MockConfig provides a test configuration
func MockConfig() *config.Config {
	return &config.Config{
		Hosts: []config.Host{
			{
				Name:       "host1",
				Address:    "192.168.1.10",
				Username:   "ubuntu",
				PrivateKey: "/home/user/.ssh/id_rsa",
			},
			{
				Name:       "host2",
				Address:    "192.168.1.11",
				Username:   "ubuntu",
				PrivateKey: "/home/user/.ssh/id_rsa",
			},
		},
		Groups: []config.HostGroup{
			{
				Name:  "all",
				Hosts: []string{"host1", "host2"},
			},
		},
	}
}

var _ ssh.Client = (*MockSSHClient)(nil)
