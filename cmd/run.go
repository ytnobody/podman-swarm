package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vibe-kanban/podman-swarm/pkg/config"
	"github.com/vibe-kanban/podman-swarm/pkg/ssh"
)

var runCmd = &cobra.Command{
	Use:   "run <host/group> <image> [args...]",
	Short: "Create and start containers",
	Long:  `Execute podman run command remotely on specified host or group.`,
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostOrGroup := args[0]
		image := args[1]
		podmanArgs := args[2:]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		hosts := cfg.GetHostOrGroup(hostOrGroup)
		if hosts == nil || len(hosts) == 0 {
			return fmt.Errorf("host or group '%s' not found", hostOrGroup)
		}

		for _, host := range hosts {
			if err := runContainerOnHost(host, image, podmanArgs); err != nil {
				fmt.Printf("Error on %s: %v\n", host.Name, err)
			}
		}
		return nil
	},
}

func runContainerOnHost(host *config.Host, image string, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ssh.NewClient(ssh.ClientConfig{
		Host:       host.Address,
		Port:       22,
		Username:   host.Username,
		PrivateKey: host.PrivateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to host: %w", err)
	}
	defer client.Close()

	cmdStr := fmt.Sprintf("podman run %s %s", image, strings.Join(args, " "))
	output, err := client.Execute(ctx, cmdStr)
	if err != nil {
		return err
	}

	fmt.Printf("[%s] %s\n", host.Name, output)
	return nil
}
