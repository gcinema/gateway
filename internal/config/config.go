// Package config
package config

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

type ServerConfig struct {
	Addr            string        `yaml:"addr"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type LoggerConfig struct {
	Level  string `yaml:"level" env-default:"debug"`
	Folder string `env:"LOGGER_FOLDER" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	path = filepath.Clean(path)
	if path == "" {
		panic("Empty path to config file")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File does not exist by path: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Cant read config")
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
