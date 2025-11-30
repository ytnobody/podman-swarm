package cmd

import (
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <host/group> <cid/name>",
	Short: "Delete containers",
	Long:  `Execute podman rm command remotely on specified host or group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement rm command
		return nil
	},
}
