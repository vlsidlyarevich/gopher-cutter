package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Path     string
	Database Database
	Server   Server
}

type Database struct {
	Url  string
	Name string
}

type Server struct {
	Port string
}

func NewConfig(path string) *Config {
	if path == "" {
		log.Panicf("Cannot instantiate config with empty file path")
	}

	return &Config{Path: path}
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile(c.Path, c); err != nil {
		log.Fatal(err)
	}
}
