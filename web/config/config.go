package config

import (
	"strings"

	"github.com/jinzhu/configor"
)

// Config is the application config
type Config struct {
	Server   Server
	Database Database
	Log      Log
}

// Log settings
type Log struct {
	Level string `default:"info"`
	Path  string
}

// Database settings
type Database struct {
	Type       string `default:"sqlite3"`
	Connection string `default:"data/ghz.db"`
}

// Server settings
type Server struct {
	RootURL string
	Port    uint `default:"3000"`
}

// Read the config file
func Read(path string) (*Config, error) {
	if strings.TrimSpace(path) == "" {
		path = "config.toml"
	}

	config := Config{}

	err := configor.Load(&config, path)

	if err != nil {
		return nil, err
	}

	return &config, nil
}