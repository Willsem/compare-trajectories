// Package config ...
package config

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func New() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
