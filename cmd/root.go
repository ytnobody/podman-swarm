package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "podman-swarm",
	Short: "Centrally manage and operate Podman container groups across multiple remote hosts via SSH",
	Long: `podman-swarm is a CLI tool for managing Podman containers across multiple remote hosts.
It uses SSH as the sole communication channel and executes native podman commands
directly on remote hosts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("podman-swarm - Podman container swarm manager")
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
	RootCmd.AddCommand(psCmd)
	RootCmd.AddCommand(inspectCmd)
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(rmCmd)
	RootCmd.AddCommand(execCmd)
}
