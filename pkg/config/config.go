package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Host struct {
	Name       string `yaml:"name"`
	Address    string `yaml:"address"`
	Username   string `yaml:"username"`
	PrivateKey string `yaml:"private_key"`
}

type HostGroup struct {
	Name  string   `yaml:"name"`
	Hosts []string `yaml:"hosts"`
}

type Config struct {
	Hosts  []Host      `yaml:"hosts"`
	Groups []HostGroup `yaml:"groups"`
}

// Load loads the configuration from the default path
func Load() (*Config, error) {
	configPath := getConfigPath()
	
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %s", configPath)
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func getConfigPath() string {
	if envPath := os.Getenv("PODMAN_SWARM_CONFIG"); envPath != "" {
		return envPath
	}
	
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "podman-swarm", "hosts.yaml")
}

// GetHostByName returns a host by its name
func (c *Config) GetHostByName(name string) *Host {
	for _, h := range c.Hosts {
		if h.Name == name {
			return &h
		}
	}
	return nil
}

// GetHostsByGroup returns all hosts in a group
func (c *Config) GetHostsByGroup(groupName string) []*Host {
	for _, g := range c.Groups {
		if g.Name == groupName {
			var hosts []*Host
			for _, hostName := range g.Hosts {
				if h := c.GetHostByName(hostName); h != nil {
					hosts = append(hosts, h)
				}
			}
			return hosts
		}
	}
	return nil
}

// GetHostOrGroup returns hosts from a name (either single host or group)
func (c *Config) GetHostOrGroup(name string) []*Host {
	if h := c.GetHostByName(name); h != nil {
		return []*Host{h}
	}
	if hosts := c.GetHostsByGroup(name); hosts != nil {
		return hosts
	}
	return nil
}
