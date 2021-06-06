package config

type ServerConfig struct {
	BindAddr string `toml:"bind_addr"`
}

type LogConfig struct {
	LogLevel string `toml:"log_level"`
}

type Config struct {
	ServerConfig `toml:"server_config"`
	LogConfig    `toml:"log_config"`
}

func New() *Config {
	return &Config{
		ServerConfig{
			BindAddr: ":8080",
		},
		LogConfig{
			LogLevel: "debug",
		},
	}
}
