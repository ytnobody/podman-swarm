## 📝 CLI Tool "podman-swarm" Requirements Definition (Final Version)

### 1. Overview and Architecture

| Item | Details |
| :--- | :--- |
| **System Name** | **podman-swarm** |
| **CLI Binary Name** | `podman-swarm` |
| **Purpose** | Using SSH as the sole communication channel, centrally manage and operate Podman container groups across multiple remote hosts from a central CLI. |
| **Architecture** | **Agent-less**. Podman-swarm (Go CLI) executes **native `podman` commands** directly via `sshd` on remote hosts. |
| **Target Container** | Podman containers |

---

### 2. Functional Requirements

#### 2.1. Inventory Management Feature (Configuration File)

Podman-swarm maintains information about managed hosts in a configuration file (recommended: YAML format).

| ID | Feature | Details |
| :--- | :--- | :--- |
| **F-01** | **Configuration File Loading** | Load host information from a YAML-formatted configuration file (default: `~/.config/podman-swarm/hosts.yaml` or similar). |
| **F-02** | **Connection Information Definition** | Be able to configure **hostname/IP address**, **SSH username**, and **SSH private key path** for each host. |
| **F-03** | **Grouping** | Define hosts in logical groups (e.g., `web`, `db`), and be able to execute commands on a group basis. |

#### 2.2. Information Collection (Reference) Features

Execute `podman` commands remotely and aggregate and format the results for display.

| ID | Command Name | Feature | Details |
| :--- | :--- | :--- | :--- |
| **F-10** | `podman-swarm status` | **Display Status of All Hosts** | Attempt SSH connections to all hosts in parallel, confirm connection availability and `podman info` execution capability, and display status in table format. |
| **F-11** | `podman-swarm ps` | **Display Container Information List for All Hosts** | Execute `podman ps -a --format json` on all hosts and aggregate results on the Manager side. Display in table format **including hostname**. |
| **F-12** | `podman-swarm inspect <host> <cid/name>` | **Display Specific Container Details** | Execute `podman inspect` on the specified host and display results in JSON format or similar. |

#### 2.3. Command (Operation) Features

Execute native `podman` commands on remote hosts for a group or single host.

| ID | Command Name | Feature | Details |
| :--- | :--- | :--- | :--- |
| **F-20** | `podman-swarm run <host/group> <image> ...` | **Create and Start Containers** | Execute `podman run` command remotely. `podman-swarm` should transparently pass subsequent arguments (`-d`, `--name`, `-p`, etc.) to the remote. |
| **F-21** | `podman-swarm stop <host/group> <cid/name>` | **Stop Containers** | Execute `podman stop` remotely. |
| **F-22** | `podman-swarm rm <host/group> <cid/name>` | **Delete Containers** | Execute `podman rm` remotely. |
| **F-23** | `podman-swarm exec <host> <cid/name> <cmd>` | **Execute Commands Inside Containers** | Execute `podman exec` remotely and return results to the Manager side. |

---

### 3. Non-Functional Requirements

| ID | Requirement | Details |
| :--- | :--- | :--- |
| **N-01** | **Development Language** | **Go (Golang)**. Should be built as a single static binary. |
| **N-02** | **Communication Protocol** | All communication should be conducted only via **SSH (Port 22)**. Do not expose other ports externally. |
| **N-03** | **Authentication Method** | Support only SSH **public key authentication**. |
| **N-04** | **Concurrent Processing** | SSH connections and command execution to multiple hosts should be processed concurrently using **goroutines** to optimize processing speed. |
| **N-05** | **Dependencies** | The runtime environment should not have dependencies on Go runtime or other languages (Python, Ruby, etc.). |
| **N-05a** | **Podman Location** | **Podman must only be installed on remote hosts being managed. The manager machine does NOT require Podman.** |
| **N-06** | **Error Handling** | Clearly notify the user of the source and details of SSH connection failures, command execution errors, JSON parsing errors, etc. |
| **N-07** | **Output Format** | Output from information display commands like `ps` should be in **formatted table format** that is easy to read in the CLI, and should also provide JSON output options (e.g., `--json`) as needed. |

---

### 4. Recommended Technology Stack

| Element | Technology | Rationale |
| :--- | :--- | :--- |
| **Language** | Go (Golang) | Lightweight binary, concurrent processing, stable SSH library. |
| **SSH Client** | `golang.org/x/crypto/ssh` | Standard Go SSH functionality. |
| **CLI Framework** | `github.com/spf13/cobra` | Standardization of command hierarchy and argument processing. |
| **Configuration/Inventory** | `github.com/spf13/viper` | YAML format configuration file loading. |
| **Table Display** | `github.com/olekukonko/tablewriter`, etc. | Format CLI output with high readability. |