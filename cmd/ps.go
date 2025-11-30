package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/vibe-kanban/podman-swarm/pkg/config"
	"github.com/vibe-kanban/podman-swarm/pkg/podman"
	"github.com/vibe-kanban/podman-swarm/pkg/ssh"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Display container information list for all hosts",
	Long:  `Execute podman ps -a on all hosts and aggregate results in table format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		jsonOutput, _ := cmd.Flags().GetBool("json")

		results := make([]*podman.ContainerListResult, 0)
		var mu sync.Mutex
		var wg sync.WaitGroup

		for _, host := range cfg.Hosts {
			wg.Add(1)
			go func(h config.Host) {
				defer wg.Done()

				result := listContainersOnHost(h)
				mu.Lock()
				results = append(results, result)
				mu.Unlock()
			}(host)
		}

		wg.Wait()

		if jsonOutput {
			displayPsJSON(results)
		} else {
			displayPsTable(results)
		}
		return nil
	},
}

func init() {
	psCmd.Flags().Bool("json", false, "Output in JSON format")
}

func listContainersOnHost(host config.Host) *podman.ContainerListResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ssh.NewClient(ssh.ClientConfig{
		Host:       host.Address,
		Port:       22,
		Username:   host.Username,
		PrivateKey: host.PrivateKey,
	})
	if err != nil {
		return &podman.ContainerListResult{
			Hostname: host.Name,
			Error:    err.Error(),
		}
	}
	defer client.Close()

	result, _ := podman.ListContainers(ctx, host.Name, client)
	return result
}

func displayPsTable(results []*podman.ContainerListResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "Container ID", "Name", "Image", "Status", "Ports"})
	table.SetBorder(true)
	table.SetRowLine(false)

	for _, result := range results {
		if result.Error != "" {
			table.Append([]string{result.Hostname, "ERROR", result.Error, "", "", ""})
			continue
		}

		for _, c := range result.Containers {
			table.Append([]string{
				result.Hostname,
				c.ID[:12],
				c.Name,
				c.Image,
				c.Status,
				c.Ports,
			})
		}
	}

	table.Render()
}

func displayPsJSON(results []*podman.ContainerListResult) {
	data, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(data))
}
