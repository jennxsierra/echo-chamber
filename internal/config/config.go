package config

import (
	"flag"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	port := flag.String("port", "4000", "Port to run the server on")
	flag.Parse()

	return Config{
		Port: *port,
	}
}
