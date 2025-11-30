package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <host/group> <image> [args...]",
	Short: "Create and start containers",
	Long:  `Execute podman run command remotely on specified host or group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement run command
		return nil
	},
}
