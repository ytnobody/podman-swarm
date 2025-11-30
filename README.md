# podman-swarm

A centralized CLI tool for managing Podman container groups across multiple remote hosts using SSH.

## Overview

`podman-swarm` allows you to manage Podman containers across multiple remote hosts from a single CLI interface. It uses SSH as the sole communication channel and executes native `podman` commands directly on remote hosts.

## Features

### Inventory Management
- Load host configurations from YAML file (`~/.config/podman-swarm/hosts.yaml`)
- Define hosts with connection details (hostname/IP, SSH username, SSH private key)
- Organize hosts into logical groups

### Information Commands
- `status` - Display status of all hosts
- `ps` - Display container information from all hosts
- `inspect` - Display specific container details

### Operation Commands
- `run` - Create and start containers
- `stop` - Stop containers
- `rm` - Delete containers
- `exec` - Execute commands inside containers

## Installation

```bash
go build -o podman-swarm
```

## Configuration

Create a configuration file at `~/.config/podman-swarm/hosts.yaml`:

```yaml
hosts:
  - name: host1
    address: 192.168.1.10
    username: ubuntu
    private_key: ~/.ssh/id_rsa
  - name: host2
    address: 192.168.1.11
    username: ubuntu
    private_key: ~/.ssh/id_rsa

groups:
  - name: web
    hosts:
      - host1
      - host2
```

## Usage

```bash
# Display status of all hosts
podman-swarm status

# List all containers across hosts
podman-swarm ps

# Inspect a specific container
podman-swarm inspect host1 container-name

# Run a container on a host or group
podman-swarm run host1 nginx:latest -d -p 80:80

# Stop a container
podman-swarm stop host1 container-name

# Remove a container
podman-swarm rm host1 container-name

# Execute a command in a container
podman-swarm exec host1 container-name /bin/sh
```

## Architecture

- **Language**: Go
- **SSH Client**: golang.org/x/crypto/ssh
- **CLI Framework**: github.com/spf13/cobra
- **Configuration**: github.com/spf13/viper
- **Table Output**: github.com/olekukonko/tablewriter

## Development

### Project Structure

```
.
├── cmd/              # CLI command implementations
├── pkg/
│   ├── config/       # Configuration file handling
│   ├── ssh/          # SSH client implementation
│   └── podman/       # Podman command wrappers
├── main.go
├── go.mod
└── go.sum
```

## Testing

```bash
go test ./...
```

## Building

```bash
go build -o podman-swarm
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute to this project.

