package main

import (
	"os"

	"github.com/ytnobody/podman-swarm/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
