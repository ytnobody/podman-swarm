package cmd

import (
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <host/group> <cid/name>",
	Short: "Stop containers",
	Long:  `Execute podman stop command remotely on specified host or group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement stop command
		return nil
	},
}
