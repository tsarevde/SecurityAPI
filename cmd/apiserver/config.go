package main

type Config struct {
	BindPort  string `toml:"bind_port"`
	Route     string `toml:"route"`
	SecretKey string `toml:"secret_key"`
}

func NewConfig() *Config {
	return &Config{
		BindPort:  ":8080",
		Route:     "/sessions",
		SecretKey: "F29Ks'iVVcB>Oqxpxklz:TQ[",
	}
}
