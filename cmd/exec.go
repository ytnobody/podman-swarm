package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/ytnobody/podman-swarm/pkg/config"
	"github.com/ytnobody/podman-swarm/pkg/ssh"
)

var execCmd = &cobra.Command{
	Use:   "exec <host> <cid/name> <cmd>",
	Short: "Execute commands inside containers",
	Long:  `Execute podman exec command remotely on specified host.`,
	Args:  cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostName := args[0]
		containerID := args[1]
		command := args[2:]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		host := cfg.GetHostByName(hostName)
		if host == nil {
			return fmt.Errorf("host '%s' not found", hostName)
		}

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

		cmdStr := fmt.Sprintf("podman exec %s %s", containerID, strings.Join(command, " "))
		output, err := client.Execute(ctx, cmdStr)
		if err != nil {
			return err
		}

		fmt.Print(output)
		return nil
	},
}
