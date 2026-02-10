package config

import "os"

type Config struct {
	Database string
}

func Load() Config {
	return Config{
		Database: os.Getenv("DATABASE"),
	}
}
