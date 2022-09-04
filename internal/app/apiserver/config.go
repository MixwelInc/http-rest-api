package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"` //server address
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// New config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
