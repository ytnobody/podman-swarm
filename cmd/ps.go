package cmd

import (
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Display container information list for all hosts",
	Long:  `Execute podman ps -a on all hosts and aggregate results in table format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement ps command
		return nil
	},
}

func init() {
	psCmd.Flags().Bool("json", false, "Output in JSON format")
}
