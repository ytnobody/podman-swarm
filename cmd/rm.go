package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/ytnobody/podman-swarm/pkg/config"
	"github.com/ytnobody/podman-swarm/pkg/ssh"
)

var rmCmd = &cobra.Command{
	Use:   "rm <host/group> <cid/name>",
	Short: "Delete containers",
	Long:  `Execute podman rm command remotely on specified host or group.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostOrGroup := args[0]
		containerID := args[1]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		hosts := cfg.GetHostOrGroup(hostOrGroup)
		if hosts == nil || len(hosts) == 0 {
			return fmt.Errorf("host or group '%s' not found", hostOrGroup)
		}

		for _, host := range hosts {
			if err := removeContainerOnHost(host, containerID); err != nil {
				fmt.Printf("Error on %s: %v\n", host.Name, err)
			}
		}
		return nil
	},
}

func removeContainerOnHost(host *config.Host, containerID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ssh.NewClient(ssh.ClientConfig{
		Host:       host.Address,
		Port:       host.Port,
		Username:   host.Username,
		PrivateKey: host.PrivateKey,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to host: %w", err)
	}
	defer client.Close()

	cmdStr := fmt.Sprintf("podman rm %s", containerID)
	output, err := client.Execute(ctx, cmdStr)
	if err != nil {
		return err
	}

	fmt.Printf("[%s] %s\n", host.Name, output)
	return nil
}
