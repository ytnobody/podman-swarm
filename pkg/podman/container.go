package podman

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
func ListContainers(hostname string, client interface{}) (*ContainerListResult, error) {
	// TODO: Implement container listing
	return &ContainerListResult{}, nil
}

// InspectContainer executes podman inspect on a remote host
func InspectContainer(hostname string, client interface{}, cid string) (map[string]interface{}, error) {
	// TODO: Implement container inspection
	return map[string]interface{}{}, nil
}
