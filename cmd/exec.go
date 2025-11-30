package cmd

import (
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec <host> <cid/name> <cmd>",
	Short: "Execute commands inside containers",
	Long:  `Execute podman exec command remotely on specified host.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement exec command
		return nil
	},
}
