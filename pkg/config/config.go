package config

type Config struct {
	ClusterAddr string
}

func NewDefaultConfig() *Config {
	return &Config{}
}
