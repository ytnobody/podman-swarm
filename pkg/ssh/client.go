package ssh

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
)

type Client interface {
	Execute(ctx context.Context, cmd string) (string, error)
	Close() error
}

type ClientConfig struct {
	Host       string
	Port       int
	Username   string
	PrivateKey string
}

type sshClient struct {
	session *ssh.Session
	client  *ssh.Client
}

// NewClient creates a new SSH client
func NewClient(config ClientConfig) (Client, error) {
	if config.Port == 0 {
		config.Port = 22
	}

	key, err := ioutil.ReadFile(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: config.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil // TODO: Implement proper host key verification
		},
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &sshClient{
		session: session,
		client:  client,
	}, nil
}

func (c *sshClient) Execute(ctx context.Context, cmd string) (string, error) {
	var stdout, stderr bytes.Buffer
	c.session.Stdout = &stdout
	c.session.Stderr = &stderr

	if err := c.session.Run(cmd); err != nil {
		return "", fmt.Errorf("command failed: %v, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

func (c *sshClient) Close() error {
	if c.session != nil {
		c.session.Close()
	}
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}
