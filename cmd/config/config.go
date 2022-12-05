package config

import (
	"os"
	"strconv"

	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/log"
	"github.com/pkg/errors"
)

// New creates a new config.
func New() (*configuration.Config, error) {
	host := os.Getenv("HOST")
	if host == "" {
		return nil, errors.New("HOST is empty")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT is empty")
	}

	var ttl int = 600

	ttlStr := os.Getenv("CACHE_TTL")
	if ttlStr != "" {
		ttl, _ = strconv.Atoi(ttlStr)
	}

	var timeout = 30

	timeoutstr := os.Getenv("TIMEOUT")
	if timeoutstr != "" {
		timeout, _ = strconv.Atoi(timeoutstr)
	}

	c := &configuration.Config{
		Log:     log.New(configuration.ServiceName, os.Getenv("LOG_LEVEL")),
		Host:    host,
		Port:    port,
		Timeout: timeout,
		TTL:     ttl,
	}

	return c, nil
}
