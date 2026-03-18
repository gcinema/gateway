package server

import (
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func NewConfig(addr string, shutdownTimeout time.Duration) Config {
	return Config{
		Addr:            addr,
		ShutdownTimeout: shutdownTimeout,
	}
}
