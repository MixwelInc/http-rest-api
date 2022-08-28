package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"` //server address
}

//New config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	} 
}