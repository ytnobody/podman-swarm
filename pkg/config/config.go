package config

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
	// TODO: Implement configuration loading
	return &Config{}, nil
}
