package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
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

	if (res == "") {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
