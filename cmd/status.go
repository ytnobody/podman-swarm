package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display status of all hosts",
	Long:  `Attempt SSH connections to all hosts in parallel and display their status in table format.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Implement status command
		return nil
	},
}
