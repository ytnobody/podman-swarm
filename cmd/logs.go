package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/ytnobody/podman-swarm/pkg/config"
	"github.com/ytnobody/podman-swarm/pkg/ssh"
)

var logsCmd = &cobra.Command{
	Use:   "logs <host/group> <cid/name>",
	Short: "Display container logs",
	Long:  `Fetch and display logs from a container on specified host.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostOrGroup := args[0]
		container := args[1]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		hosts := cfg.GetHostOrGroup(hostOrGroup)
		if hosts == nil || len(hosts) == 0 {
			return fmt.Errorf("host or group '%s' not found", hostOrGroup)
		}

		for _, host := range hosts {
			if err := getContainerLogs(host, container); err != nil {
				fmt.Printf("Error on %s: %v\n", host.Name, err)
			}
		}
		return nil
	},
}

func getContainerLogs(host *config.Host, container string) error {
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

	cmdStr := fmt.Sprintf("podman logs %s", container)
	output, err := client.Execute(ctx, cmdStr)
	if err != nil {
		return err
	}

	fmt.Printf("[%s]\n%s\n", host.Name, output)
	return nil
}

func init() {
	RootCmd.AddCommand(logsCmd)
}
