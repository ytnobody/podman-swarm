package cmd

import (
	"context"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/ytnobody/podman-swarm/pkg/config"
	"github.com/ytnobody/podman-swarm/pkg/ssh"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display status of all hosts",
	Long:  `Attempt SSH connections to all hosts in parallel and display their status in table format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		results := make([]statusResult, 0)
		var mu sync.Mutex
		var wg sync.WaitGroup

		for _, host := range cfg.Hosts {
			wg.Add(1)
			go func(h config.Host) {
				defer wg.Done()

				result := checkHostStatus(h)
				mu.Lock()
				results = append(results, result)
				mu.Unlock()
			}(host)
		}

		wg.Wait()

		displayStatusTable(results)
		return nil
	},
}

type statusResult struct {
	Host   string
	Status string
	Error  string
}

func checkHostStatus(host config.Host) statusResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ssh.NewClient(ssh.ClientConfig{
		Host:       host.Address,
		Port:       host.Port,
		Username:   host.Username,
		PrivateKey: host.PrivateKey,
	})
	if err != nil {
		return statusResult{
			Host:   host.Name,
			Status: "DOWN",
			Error:  err.Error(),
		}
	}
	defer client.Close()

	_, err = client.Execute(ctx, "podman info > /dev/null 2>&1")
	if err != nil {
		return statusResult{
			Host:   host.Name,
			Status: "DOWN",
			Error:  err.Error(),
		}
	}

	return statusResult{
		Host:   host.Name,
		Status: "UP",
	}
}

func displayStatusTable(results []statusResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "Status", "Details"})
	table.SetBorder(true)
	table.SetRowLine(false)

	for _, r := range results {
		details := r.Error
		if details == "" {
			details = "OK"
		}
		table.Append([]string{r.Host, r.Status, details})
	}

	table.Render()
}
