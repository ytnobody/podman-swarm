package podman

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ytnobody/podman-swarm/pkg/ssh"
)

type Container struct {
	ID      string
	Name    string
	Image   string
	Status  string
	Ports   string
	Created string
}

type ContainerListResult struct {
	Hostname   string
	Containers []Container
	Error      string
}

// ListContainers executes podman ps on a remote host
func ListContainers(ctx context.Context, hostname string, client ssh.Client) (*ContainerListResult, error) {
	result := &ContainerListResult{Hostname: hostname}

	output, err := client.Execute(ctx, "podman ps -a --format json")
	if err != nil {
		result.Error = err.Error()
		return result, nil
	}

	var containers []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &containers); err != nil {
		result.Error = fmt.Sprintf("failed to parse JSON: %v", err)
		return result, nil
	}

	for _, c := range containers {
		container := Container{
			ID:      toString(c["Id"]),
			Name:    toString(c["Names"]),
			Image:   toString(c["Image"]),
			Status:  toString(c["Status"]),
			Ports:   toString(c["Ports"]),
			Created: toString(c["Created"]),
		}
		result.Containers = append(result.Containers, container)
	}

	return result, nil
}

// InspectContainer executes podman inspect on a remote host
func InspectContainer(ctx context.Context, hostname string, client ssh.Client, cid string) (map[string]interface{}, error) {
	output, err := client.Execute(ctx, fmt.Sprintf("podman inspect %s", cid))
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("container not found")
	}

	return result[0], nil
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case []interface{}:
		if len(val) > 0 {
			return toString(val[0])
		}
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}
