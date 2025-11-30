package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/vibe-kanban/podman-swarm/pkg/config"
	"github.com/vibe-kanban/podman-swarm/pkg/podman"
	"github.com/vibe-kanban/podman-swarm/pkg/ssh"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect <host> <cid/name>",
	Short: "Display specific container details",
	Long:  `Execute podman inspect on the specified host and display results in JSON format.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostName := args[0]
		containerID := args[1]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		host := cfg.GetHostByName(hostName)
		if host == nil {
			return fmt.Errorf("host '%s' not found in configuration", hostName)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

		result, err := podman.InspectContainer(ctx, hostName, client, containerID)
		if err != nil {
			return err
		}

		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
		return nil
	},
}
