package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Host struct {
	Name       string `mapstructure:"name" yaml:"name"`
	Address    string `mapstructure:"address" yaml:"address"`
	Port       int    `mapstructure:"port" yaml:"port"`
	Username   string `mapstructure:"username" yaml:"username"`
	PrivateKey string `mapstructure:"private_key" yaml:"private_key"`
}

type HostGroup struct {
	Name  string   `mapstructure:"name" yaml:"name"`
	Hosts []string `mapstructure:"hosts" yaml:"hosts"`
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

	for i := range cfg.Hosts {
		if cfg.Hosts[i].PrivateKey == "" {
			return nil, fmt.Errorf("private_key is empty for host %s", cfg.Hosts[i].Name)
		}
		cfg.Hosts[i].PrivateKey = expandPath(cfg.Hosts[i].PrivateKey)
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

func expandPath(path string) string {
	if path == "" {
		return path
	}
	if path[0] == '~' {
		homeDir, _ := os.UserHomeDir()
		if len(path) > 1 && path[1] == '/' {
			return filepath.Join(homeDir, path[2:])
		}
		return homeDir
	}
	return path
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
