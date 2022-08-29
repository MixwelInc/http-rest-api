package apiserver

import "github.com/MixwelInc/http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` //server address
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// New config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
