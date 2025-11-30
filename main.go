package main

import (
	"os"

	"github.com/vibe-kanban/podman-swarm/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
