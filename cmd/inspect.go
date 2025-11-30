package cmd

import (
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect <host> <cid/name>",
	Short: "Display specific container details",
	Long:  `Execute podman inspect on the specified host and display results in JSON format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement inspect command
		return nil
	},
}
