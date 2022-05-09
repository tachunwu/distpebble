package config

type Config struct {
	SequencerAddr string
	ClusterAddr   []string
}

func NewDefaultConfig() *Config {
	return &Config{}
}
